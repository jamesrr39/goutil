package gofs

import (
	"io/fs"
	"os"
	"path/filepath"
	"strings"
	"sync"

	"github.com/jamesrr39/goutil/patternmatcher"
)

type WalkOptions struct {
	IncludesMatcher,
	ExcludesMatcher patternmatcher.Matcher
	MaxConcurrency uint
	// FollowSymlinks treats a symlink as if it were the file/directory underneath, effectively hiding the symlink from the consumer.
	// It can be used with a chain of symlinks, in which case it will keep following the links, until it reaches a non-symlink
	FollowSymlinks bool
}

type walkerType struct {
	fs              Fs
	basePath        string
	walkFunc        filepath.WalkFunc
	options         WalkOptions
	errChan         chan error
	addToQueueWg    *sync.WaitGroup
	processPathChan chan string
}

const DefaultMaxConcurrency = 1

// Walk walks a tree concurrently
// If options.MaxConcurrency is set to >1, any actions you do in the provided walkFunc must be synchronized (or not need to be synchronized)
func Walk(fs Fs, path string, walkFunc filepath.WalkFunc, options WalkOptions) error {
	maxConcurrency := options.MaxConcurrency
	if maxConcurrency == 0 {
		maxConcurrency = DefaultMaxConcurrency
	}

	wt := &walkerType{
		fs:              fs,
		basePath:        path,
		walkFunc:        walkFunc,
		options:         options,
		errChan:         make(chan error),
		addToQueueWg:    new(sync.WaitGroup),
		processPathChan: make(chan string, maxConcurrency),
	}

	doneChan := make(chan error)

	go func() {
		for {
			select {
			case err := <-wt.errChan:
				doneChan <- err
				return

			case path := <-wt.processPathChan:
				go func(path string) {
					defer wt.addToQueueWg.Done()
					fileInfo, err := wt.processPath(path)
					if err != nil {
						wt.errChan <- err
						return
					}

					if fileInfo == nil {
						return
					}

					if fileInfo.IsDir() {
						err = wt.walkDir(path)
						if err != nil {
							wt.errChan <- err
							return
						}
					}
				}(path)
			}
		}
	}()

	fileInfo, err := wt.fs.Lstat(path)
	if err != nil {
		return err
	}

	err = wt.addToProcessPathChan(path, fileInfo)
	if err != nil {
		return err
	}

	go func() {
		wt.addToQueueWg.Wait()
		doneChan <- nil
	}()

	err = <-doneChan

	return err
}

func (wt *walkerType) processPath(path string) (os.FileInfo, error) {
	relativePath := strings.TrimPrefix(strings.TrimPrefix(path, wt.basePath), string(filepath.Separator))
	isExcluded := wt.options.ExcludesMatcher != nil && wt.options.ExcludesMatcher.Matches(relativePath)
	if isExcluded {
		return nil, nil
	}

	if wt.options.IncludesMatcher != nil {
		isIncluded := wt.options.IncludesMatcher.Matches(relativePath)
		if !isIncluded {
			return nil, nil
		}
	}

	fileInfo, err := wt.fs.Lstat(path)
	if err != nil {
		return nil, err
	}

	err = wt.walkFunc(path, fileInfo, nil)
	if err != nil {
		return nil, err
	}

	return fileInfo, nil
}

func (wt *walkerType) walkDir(path string) error {
	dirEntryInfos, err := wt.fs.ReadDir(path)
	if err != nil {
		return err
	}

	for _, dirEntryInfo := range dirEntryInfos {
		childPath := filepath.Join(path, dirEntryInfo.Name())

		wt.addToProcessPathChan(childPath, dirEntryInfo)
	}

	return nil
}

func (wt *walkerType) addToProcessPathChan(path string, fileInfo fs.FileInfo) error {
	if wt.options.FollowSymlinks {
		targetPath, _, err := wt.resolveSymlink(path, fileInfo)
		if err != nil {
			return err
		}

		path = targetPath
	}

	wt.addToQueueWg.Add(1)
	wt.processPathChan <- path
	return nil
}

func (wt *walkerType) resolveSymlink(path string, fileInfo os.FileInfo) (string, os.FileInfo, error) {
	for IsSymlink(fileInfo.Mode()) {
		targetPath, err := wt.fs.Readlink(path)
		if err != nil {
			return "", nil, err
		}

		// since Walk() takes in a path, if the symlink is a relative path,
		// we need to turn that into a path relative from the path that was passed in to Walk()
		if !filepath.IsAbs(targetPath) {
			dirPath := filepath.Dir(path)
			targetPath = filepath.Join(dirPath, targetPath)
		}

		targetInfo, err := wt.fs.Lstat(targetPath)
		if err != nil {
			return "", nil, err
		}

		path = targetPath
		fileInfo = targetInfo
	}

	return path, fileInfo, nil
}
