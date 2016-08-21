package rnm

import (
	"os"
)

type fileRenamer interface {
	Exists(path string) bool
	Rename(oldPath string, newPath string) error
}

type actualRenamer struct{}

func (actualRenamer) Exists(path string) bool {
	_, err := os.Stat(path)
	return err == nil
}

func (actualRenamer) Rename(oldPath string, newPath string) error {
	return os.Rename(oldPath, newPath)
}
