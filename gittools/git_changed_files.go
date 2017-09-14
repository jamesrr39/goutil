package gittools

import (
	"bufio"
	"bytes"
	"os/exec"
	"strings"
)

type FileChangeType int

const (
	Added FileChangeType = iota
	Modified
	Removed
	Other
)

type FileChange struct {
	FilePath   string // relative to git base
	ChangeType FileChangeType
}

/*
 M b.ts
?? a.ts
*/
func GetChangedFiles(repoPath string) ([]*FileChange, error) {
	cmd := exec.Command("git", "-C", repoPath, "status", "--porcelain") // todo could be "-C repoPath"
	outputBytes, err := cmd.Output()
	if nil != err {
		return nil, err
	}

	var changedFiles []*FileChange
	s := bufio.NewScanner(bytes.NewBuffer(outputBytes))
	for s.Scan() {
		fields := strings.Fields(s.Text())
		changedFile := &FileChange{
			fields[1],
			fileChangeTypeFromGitString(fields[0]),
		}
		changedFiles = append(changedFiles, changedFile)
	}

	return changedFiles, nil
}

func fileChangeTypeFromGitString(gitString string) FileChangeType {
	switch gitString {
	case "M":
		return Modified
	default:
		return Other
	}
}
