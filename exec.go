package rnm

import (
	"path/filepath"
)

type Result struct {
	OldPath string
	NewPath string
	Error   error
}

func Exec(patterns []string, opts Option) (results []Result, err error) {
	if opts.From == opts.To {
		return []Result{}, nil
	}

	targetPaths, err := listFiles(zGlobber{}, patterns, opts)

	if err != nil {
		return nil, err
	}

	results = make([]Result, len(targetPaths))
	renamer := actualRenamer{}

	for i, path := range targetPaths {
		dirPath, fileName := filepath.Split(path)

		newName := convert(fileName, convertOption{
			From:     opts.From,
			To:       opts.To,
			AsRegexp: opts.AsRegexp,
		})

		oldPath := dirPath + fileName
		newPath := dirPath + newName

		err := renameFile(renamer, renameOption{
			OldPath: oldPath,
			NewPath: newPath,
			Dryrun:  opts.Dryrun,
		})

		results[i] = Result{
			OldPath: oldPath,
			NewPath: newPath,
			Error:   err,
		}
	}

	return results, nil
}
