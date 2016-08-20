package rnm

type fileRenamer interface {
	Exists(path string) bool
	Rename(oldPath string, newPath string) error
}

type actualRenamer struct{}

func (actualRenamer) Exists(path string) bool {
	return true
}

func (actualRenamer) Rename(oldPath string, newPath string) error {
	return nil
}
