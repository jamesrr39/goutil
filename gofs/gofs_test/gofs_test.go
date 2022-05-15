package gofstest

import (
	"fmt"
	"io/fs"
	"sync/atomic"
	"testing"
	"time"

	"github.com/jamesrr39/goutil/gofs"
	"github.com/jamesrr39/goutil/gofs/mockfs"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func Test_concurrency(t *testing.T) {
	const pathsToMake = 10000
	mockFs := mockfs.NewMockFs()
	mockFs.LstatFunc = mockFs.Stat

	dirPaths := []string{"/dir1", "/dir2"}

	for _, dirPath := range dirPaths {
		err := mockFs.MkdirAll(dirPath, 0644)
		require.NoError(t, err)
		for i := 0; i < pathsToMake; i++ {
			_, err := mockFs.Create(fmt.Sprintf("%s/file_%d.txt", dirPath, i))
			require.NoError(t, err)
		}
	}

	var totalPathsScanned int64

	walkFunc := func(path string, info fs.FileInfo, err error) error {
		time.Sleep(time.Second / 2)
		atomic.AddInt64(&totalPathsScanned, 1)
		return nil
	}

	opts := gofs.WalkOptions{
		MaxConcurrency: 1000,
	}

	err := gofs.Walk(mockFs, "/", walkFunc, opts)
	require.NoError(t, err)

	const expectedPathsToBeScanned = pathsToMake*2 + 3 // each file under dir1/ and dir2/, then the root dir, dir1 and dir2 entries
	assert.Equal(t, pathsToMake, totalPathsScanned)
}
