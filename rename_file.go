package rnm

type NewPathExistError struct{}

func (e *NewPathExistError) Error() string {
	return "The new path already exists"
}

func renameFile(renamer fileRenamer, opts *renameOption) error {
	oldPath, newPath := opts.OldPath, opts.NewPath

	if renamer.Exists(newPath) {
		return new(NewPathExistError)
	}

	if opts.Dryrun {
		return nil
	}

	err := renamer.Rename(oldPath, newPath)
	return err
}
