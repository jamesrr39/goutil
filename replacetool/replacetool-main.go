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
	fmt.Printf(`a script for replacing "return err" with a custom return.

Usage:

replacetool path/to/project errorsx.Wrap(err) github.com/jamesrr39/goutil/errorsx

Where
- "path/to/project" is the path to your project that you want the replaces to happen inside
- "errorsx.Wrap(err)" is the new string to replace with
- "github.com/jamesrr39/goutil/errorsx" is the package name required for the new string
`)
}

func main() {
	shouldReplace := flag.Bool("replace", false, "should replace err with errors.Wrap(err)")
	flag.Usage = usage
	flag.Parse()

	if len(flag.Args()) < 3 {
		flag.Usage()
		os.Exit(1)
	}
	basePath := flag.Arg(0)
	replaceString := flag.Arg(1)
	errorsPkgImportPath := flag.Arg(2)

	replacer := &Replacer{
		basePath:            basePath,
		shouldReplace:       *shouldReplace,
		replaceString:       replaceString,
		errorsPkgImportPath: errorsPkgImportPath,
	}

	allSuccess, err := replacer.run()
	if err != nil {
		log.Fatalln(err)
	}

	if !allSuccess {
		os.Exit(1)
	}
}

type Replacer struct {
	basePath            string
	shouldReplace       bool
	replaceString       string
	errorsPkgImportPath string
}

func (r *Replacer) run() (allSuccess bool, err error) {
	allSuccess = true
	err = filepath.Walk(r.basePath, func(path string, fileInfo os.FileInfo, err error) error {
		if err != nil {
			return err
		}

		if fileInfo.IsDir() {
			return nil
		}

		if !strings.HasSuffix(fileInfo.Name(), ".go") {
			return nil
		}

		relativePath := strings.TrimPrefix(strings.TrimPrefix(path, r.basePath), string(filepath.Separator))

		if strings.HasPrefix(relativePath, "vendor"+string(filepath.Separator)) || strings.HasPrefix(relativePath, ".git"+string(filepath.Separator)) {
			return nil
		}

		ok, err := r.runFile(path, fileInfo, r.shouldReplace)
		if err != nil {
			return err
		}

		if !ok {
			allSuccess = false
		}

		return nil
	})
	return allSuccess, err
}

func (r *Replacer) runFile(filePath string, fileInfo os.FileInfo, replace bool) (bool, error) {
	fileBytes, err := ioutil.ReadFile(filePath)
	if err != nil {
		return false, err
	}
	sourceCode := string(fileBytes)

	parsedFile, err := parser.ParseFile(token.NewFileSet(), "", sourceCode, 0)
	if err != nil {
		return false, err
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

	if len(nodes) == 0 {
		return true, nil
	}

	if replace {
		return true, r.replaceFileContents(filePath, fileInfo, sourceCode, nodes, parsedFile)
	}
	for _, node := range nodes {
		fmt.Printf("found unwrapped err in %s (pos %d-%d)\n", filePath, node.Pos(), node.End())
	}
	return false, nil
}

func (r *Replacer) replaceFileContents(filePath string, fileInfo os.FileInfo, sourceCode string, nodes []ast.Expr, parsedFile *ast.File) error {

	for i := len(nodes) - 1; i >= 0; i-- {
		node := nodes[i]
		sourceCode = fmt.Sprintf("%s%s%s", sourceCode[:node.Pos()-1], r.replaceString, sourceCode[node.End()-1:])
	}

	containsPkgErrors := false
	for _, importStmt := range parsedFile.Imports {
		if importStmt.Path.Value == fmt.Sprintf(`"%s"`, r.errorsPkgImportPath) {
			containsPkgErrors = true
		}
	}

	if !containsPkgErrors {
		if len(parsedFile.Imports) == 0 {
			return fmt.Errorf("TODO: no imports in %s, please add %q import by hand", filePath, r.errorsPkgImportPath)
		}

		lastImport := parsedFile.Imports[len(parsedFile.Imports)-1]
		sourceCode = fmt.Sprintf(`%s
	"%s"
%s`, sourceCode[:lastImport.End()], r.errorsPkgImportPath, sourceCode[lastImport.End():])
	}

	err := ioutil.WriteFile(filePath, []byte(sourceCode), fileInfo.Mode())
	if err != nil {
		return err
	}

	return nil
}
