package fswalker

import (
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"sort"
	"strings"
)

type Fs interface {
	Stat(path string) (os.FileInfo, error)
	Readlink(path string) (string, error)
	ReadDir(path string) ([]os.FileInfo, error)
}

type OsFs struct{}

func (fs OsFs) Stat(path string) (os.FileInfo, error) {
	return os.Stat(path)
}

func (fs OsFs) Readlink(path string) (string, error) {
	return os.Readlink(path)
}

func (fs OsFs) ReadDir(path string) ([]os.FileInfo, error) {
	return ioutil.ReadDir(path)
}

type ExcludeMatcher interface {
	Matches(path string) bool
	MatchesDir(path string) bool
}

type WalkOptions struct {
	FollowSymlinks bool
	Fs             Fs
	ExcludeMatcher ExcludeMatcher
}

func Walk(path string, walkFunc filepath.WalkFunc, options WalkOptions) error {
	if options.Fs == nil {
		options.Fs = OsFs{}
	}

	return walk(path, path, walkFunc, options)
}

func walk(rootPath, path string, walkFunc filepath.WalkFunc, options WalkOptions) error {
	relativePath := strings.TrimPrefix(strings.TrimPrefix(path, rootPath), string(filepath.Separator))

	if options.ExcludeMatcher.Matches(relativePath) {
		log.Printf("skipping %q\n", path) // TODO remove
		return nil
	}

	fileInfo, err := options.Fs.Stat(path)
	if err != nil {
		log.Printf("skipping bad stat of %q. Error: %q\n", path, err)
		return nil // TODO: better resolution
	}

	if fileInfo.IsDir() {
		if options.ExcludeMatcher.MatchesDir(relativePath) {
			log.Printf("skipping dir: %q\n", path) // TODO remove
			return nil
		}

		dirEntryInfos, err := options.Fs.ReadDir(path)
		if err != nil {
			return err
		}

		sort.Slice(dirEntryInfos, func(i int, j int) bool {
			return dirEntryInfos[i].Name() > dirEntryInfos[j].Name()
		})

		for _, dirEntryInfo := range dirEntryInfos {
			err = walk(rootPath, filepath.Join(path, dirEntryInfo.Name()), walkFunc, options)
			if err != nil {
				return err
			}
		}
	}

	if options.FollowSymlinks {
		isSymlink := (fileInfo.Mode() & os.ModeSymlink) == 1
		if isSymlink {
			linkDest, err := options.Fs.Readlink(path)
			if err != nil {
				return err
			}
			err = walk(rootPath, linkDest, walkFunc, options)
			if err != nil {
				return err
			}
		}
	}

	err = walkFunc(path, fileInfo, nil)
	if err != nil {
		return err
	}

	return nil
}
