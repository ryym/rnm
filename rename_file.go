package rnm

import (
	"fmt"
)

type NewPathExistError struct {
	NewPath string
}

func (e *NewPathExistError) Error() string {
	return fmt.Sprintf("The new path already exists: %s", e.NewPath)
}

func renameFile(renamer fileRenamer, opts *renameOption) error {
	oldPath, newPath := opts.OldPath, opts.NewPath

	if renamer.Exists(newPath) {
		return &NewPathExistError{newPath}
	}

	if opts.Dryrun {
		return nil
	}

	err := renamer.Rename(oldPath, newPath)
	return err
}
