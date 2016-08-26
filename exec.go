package rnm

import (
	"path/filepath"
	"sort"
	"strings"
)

type Result struct {
	OldPath string
	NewPath string
	Error   error
}

type targetPaths []string

func (paths targetPaths) Len() int {
	return len(paths)
}

func (paths targetPaths) Swap(i, j int) {
	paths[i], paths[j] = paths[j], paths[i]
}

// Ensure a directory is renamed after the files
// it has are renamed:
//   pic1 -> photo1
//   pic2 -> photo2
//   pics/pic3 -> pics/photo3
//   pics/pic4 -> pics/photo4
//   pics -> photos
func (paths targetPaths) Less(i, j int) bool {
	if strings.Contains(paths[i], paths[j]) {
		return paths[i] > paths[j]
	} else {
		return paths[i] < paths[j]
	}
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

	var paths targetPaths = selectTargetPaths(converter, candidates)
	sort.Sort(paths)

	results = make([]Result, len(paths))
	renamer := actualRenamer{}

	for i, path := range paths {
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
