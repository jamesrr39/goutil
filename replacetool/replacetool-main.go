package main

import (
	"flag"
	"fmt"
	"go/ast"
	"go/parser"
	"go/token"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"strings"
)

var usage = func() {
	fmt.Println(`a script for replacing "return err" with return errors.WithStack(err)`)
}

func main() {
	shouldReplace := flag.Bool("replace", false, "should replace err with errors.Wrap(err)")
	flag.Usage = usage
	flag.Parse()
	basePath := flag.Arg(0)

	err := run(basePath, *shouldReplace)
	if err != nil {
		log.Fatalln(err)
	}
}

func run(basePath string, shouldReplace bool) error {
	return filepath.Walk(basePath, func(path string, fileInfo os.FileInfo, err error) error {
		if err != nil {
			return err
		}

		if fileInfo.IsDir() {
			return nil
		}

		if !strings.HasSuffix(fileInfo.Name(), ".go") {
			return nil
		}

		relativePath := strings.TrimPrefix(strings.TrimPrefix(path, basePath), string(filepath.Separator))

		if strings.HasPrefix(relativePath, "vendor"+string(filepath.Separator)) || strings.HasPrefix(relativePath, ".git"+string(filepath.Separator)) {
			return nil
		}

		return runFile(path, fileInfo, shouldReplace)
	})
}

func runFile(filePath string, fileInfo os.FileInfo, replace bool) error {
	fileBytes, err := ioutil.ReadFile(filePath)
	if err != nil {
		return err
	}
	sourceCode := string(fileBytes)

	parsedFile, err := parser.ParseFile(token.NewFileSet(), "", sourceCode, 0)
	if err != nil {
		return err
	}

	var nodes []ast.Expr

	ast.Inspect(parsedFile, func(node ast.Node) bool {
		switch n := node.(type) {
		case *ast.ReturnStmt:
			for _, returnStmt := range n.Results {
				returnVarName := sourceCode[returnStmt.Pos()-1 : returnStmt.End()-1]
				if returnVarName == "err" {
					nodes = append(nodes, returnStmt)
				}
			}

		}
		return true
	})

	if replace && len(nodes) > 0 {
		for i := len(nodes) - 1; i >= 0; i-- {
			node := nodes[i]
			sourceCode = fmt.Sprintf("%serrors.WithStack(err)%s", sourceCode[:node.Pos()-1], sourceCode[node.End()-1:])
		}

		containsPkgErrors := false
		for _, importStmt := range parsedFile.Imports {
			if importStmt.Path.Value == `"github.com/pkg/errors"` {
				containsPkgErrors = true
			}
		}

		if !containsPkgErrors {
			if len(parsedFile.Imports) == 0 {
				return fmt.Errorf("TODO: no imports in %s, please add github.com/pkg/errors import by hand", filePath)
			}

			lastImport := parsedFile.Imports[len(parsedFile.Imports)-1]
			sourceCode = fmt.Sprintf(`%s
	"github.com/pkg/errors"
%s`, sourceCode[:lastImport.End()], sourceCode[lastImport.End():])
		}

		err = ioutil.WriteFile(filePath, []byte(sourceCode), fileInfo.Mode())
		if err != nil {
			return err
		}
	}

	return nil
}
