package gofs

import (
	"io/ioutil"
	"os"
	"path/filepath"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func Test_walk(t *testing.T) {
	tmpDir, err := ioutil.TempDir("", "")
	require.NoError(t, err)
	defer os.RemoveAll(tmpDir)

	/*
		directory structure:

		- (root dir)
		  - dir_a
		    - file_a.txt
		    - symlink_to_dir_b
		  - dir_b
		    - file_b.txt
	*/

	dirAPath := filepath.Join(tmpDir, "dir_a")
	dirBPath := filepath.Join(tmpDir, "dir_b")
	fileAPath := filepath.Join(dirAPath, "file_a.txt")
	symlinkToDirBPath := filepath.Join(tmpDir, "dir_a", "symlink_to_dir_b")
	fileBPath := filepath.Join(dirBPath, "file_b.txt")
	textA := []byte("text_a")
	textB := []byte("text_b")

	err = os.Mkdir(dirAPath, 0700)
	require.NoError(t, err)

	err = os.Mkdir(dirBPath, 0700)
	require.NoError(t, err)

	err = os.Symlink(
		dirBPath,
		symlinkToDirBPath,
	)
	require.NoError(t, err)

	err = ioutil.WriteFile(fileAPath, textA, 0600)
	require.NoError(t, err)

	err = ioutil.WriteFile(fileBPath, textB, 0600)
	require.NoError(t, err)

	t.Run("FollowSymlinks: false", func(t *testing.T) {
		var visitedFileA, visitedFileB bool
		err := Walk(NewOsFs(), dirAPath, func(path string, info os.FileInfo, err error) error {
			if err != nil {
				return err
			}

			switch path {
			case fileAPath:
				visitedFileA = true
			case fileBPath:
				visitedFileB = true
			}
			return nil
		}, WalkOptions{
			MaxConcurrency: 1000,
		})
		require.NoError(t, err)

		assert.True(t, visitedFileA)
		assert.False(t, visitedFileB)
	})
}
