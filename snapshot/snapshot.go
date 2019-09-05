package snapshot

import (
	"bytes"
	"encoding/json"
	"io"
	"net/url"
	"os"
	"path/filepath"
	"runtime"
	"testing"
)

type snapshotType struct {
	Value string
}

func AssertMatchesSnapshot(t *testing.T, snapshotName, actual string) {
	if snapshotName == "" {
		t.Error("no snapshot name provided")
	}

	_, filePath, _, ok := runtime.Caller(1)
	if !ok {
		panic("couldn't call runtime.Caller")
	}

	// encode snapshot name to escape "/" and other filesystem-illegal characters
	encodedSnapshotName := url.PathEscape(snapshotName)

	snapshotDirPath := filepath.Join(filepath.Dir(filePath), "__snapshots__", filepath.Base(filePath))
	snapshotFilePath := filepath.Join(snapshotDirPath, encodedSnapshotName+".snap.json")

	actioner := shouldUpdateViaEnv()

	fileSuccessfullyOpened := true
	file, err := os.Open(snapshotFilePath)
	if err != nil {
		fileSuccessfullyOpened = false
		actioner.OnSnapshotFileOpenError(t, snapshotDirPath, snapshotFilePath, actual)
	}
	defer file.Close()

	if !fileSuccessfullyOpened {
		return
	}

	snapshot := new(snapshotType)
	err = json.NewDecoder(file).Decode(&snapshot)
	if err != nil {
		panic(err)
	}

	if snapshot.Value != actual {
		actioner.OnSnapshotNotMatched(t, file, snapshot.Value, actual)
	}
}

func shouldUpdateViaEnv() snapshotActioner {
	shouldUpdateVal, ok := os.LookupEnv("UPDATE_SNAPSHOTS")
	if ok && shouldUpdateVal == "1" {
		return updateSnapshotActioner{}
	}

	return noUpdateSnapshotActioner{}
}

type snapshotActioner interface {
	OnSnapshotFileOpenError(t *testing.T, snapshotDirPath, snapshotFilePath, actual string)
	OnSnapshotNotMatched(t *testing.T, file io.Writer, snapshotValue, actual string)
}

type noUpdateSnapshotActioner struct {
}

func (a noUpdateSnapshotActioner) OnSnapshotFileOpenError(t *testing.T, snapshotDirPath, snapshotFilePath, actual string) {
	t.Errorf("couldn't open snapshot file at %q", snapshotFilePath)
}

func (a noUpdateSnapshotActioner) OnSnapshotNotMatched(t *testing.T, file io.Writer, snapshotValue, actual string) {
	t.Errorf("expected %q but got %q", snapshotValue, actual)
}

type updateSnapshotActioner struct {
}

func (a updateSnapshotActioner) OnSnapshotFileOpenError(t *testing.T, snapshotDirPath, snapshotFilePath, actual string) {
	err := os.MkdirAll(snapshotDirPath, 0755)
	if err != nil {
		panic(err)
	}
	file, err := os.Create(snapshotFilePath)
	if err != nil {
		panic(err)
	}

	newSnapshot := &snapshotType{actual}
	b, err := json.MarshalIndent(newSnapshot, "", "\t")
	if err != nil {
		panic(err)
	}

	_, err = io.Copy(file, bytes.NewBuffer(b))
	if err != nil {
		panic(err)
	}
}

func (a updateSnapshotActioner) OnSnapshotNotMatched(t *testing.T, file io.Writer, snapshotValue, actual string) {
	newSnapshot := &snapshotType{actual}
	b, err := json.MarshalIndent(newSnapshot, "", "\t")
	if err != nil {
		panic(err)
	}

	_, err = io.Copy(file, bytes.NewBuffer(b))
	if err != nil {
		panic(err)
	}
}
