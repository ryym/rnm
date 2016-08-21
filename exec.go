package rnm

import (
	"path/filepath"
)

type result struct {
	OldPath string
	NewPath string
	Error   error
}

func Exec(patterns []string, opts Option) (results []result, err error) {
	if opts.From == opts.To {
		return []result{}, nil
	}

	targetPaths, err := listFiles(patterns, opts)

	if err != nil {
		return nil, err
	}

	results = make([]result, len(targetPaths))
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

		results[i] = result{
			OldPath: oldPath,
			NewPath: newPath,
			Error:   err,
		}
	}

	return results, nil
}
