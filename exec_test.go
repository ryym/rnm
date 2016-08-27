package rnm

import (
	"io/ioutil"
	"os"
	"path/filepath"
	"testing"
)

func TestExec(t *testing.T) {
	tmpdir, err := prepareFiles("foo", "a.txt", "b.txt", "c.txt")
	if tmpdir != "" {
		defer os.RemoveAll(tmpdir)
	}
	if err != nil {
		t.Fatal(err)
	}

	testDir := filepath.Join(tmpdir, "foo")

	Exec(
		[]string{filepath.Join(testDir, "*.txt")},
		&Option{
			From: ".txt",
			To:   "-new.txt",
		},
	)

	names, err := getFileNamesIn(testDir)
	if err != nil {
		t.Fatal(err)
	}

	expected := []string{"a-new.txt", "b-new.txt", "c-new.txt"}
	if !isSameSlices(names, expected) {
		t.Errorf(
			"Files don't be renamed. expected: %v, got: %v",
			expected,
			names,
		)
	}
}

func prepareFiles(dir string, fileNames ...string) (tmpdir string, err error) {
	tmpdir, err = ioutil.TempDir("", "rnm")
	if err != nil {
		return "", err
	}

	testDir := filepath.Join(tmpdir, dir)

	os.Mkdir(testDir, 0755)
	for _, name := range fileNames {
		err = ioutil.WriteFile(filepath.Join(testDir, name), []byte{}, 0644)
		if err != nil {
			return "", err
		}
	}

	return tmpdir, nil
}

func getFileNamesIn(dir string) (names []string, err error) {
	files, err := ioutil.ReadDir(dir)
	if err != nil {
		return nil, err
	}

	names = make([]string, len(files))
	for i, file := range files {
		names[i] = file.Name()
	}

	return names, nil
}
