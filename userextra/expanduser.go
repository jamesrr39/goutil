package userextra

import (
	"os/user"
	"strings"
)

// expand tilde.
// Returns the resolved path, or an error from trying to resolve the current user
func ExpandUser(path string) (string, error) {
	if !strings.HasPrefix(path, "~/") {
		return path, nil
	}

	u, err := user.Current()
	if nil != err {
		return "", err
	}

	return strings.Replace(path, "~", u.HomeDir, 1), nil
}
