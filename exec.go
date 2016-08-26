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

	candidates, err := listCandidates(zGlobber{}, patterns)
	if err != nil {
		return nil, err
	}

	converter, err := createConverter(opts)
	if err != nil {
		return nil, err
	}

	targetPaths := selectTargetPaths(converter, candidates)

	results = make([]Result, len(targetPaths))
	renamer := actualRenamer{}

	for i, path := range targetPaths {
		dirPath, fileName := filepath.Split(path)
		newName := converter.convert(fileName)

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

func createConverter(opts Option) (converter converter, err error) {
	copts := convertOption{
		From: opts.From,
		To:   opts.To,
	}
	if opts.AsRegexp {
		return newRegexpConverter(copts)
	} else {
		return newStringConverter(copts), nil
	}
}
