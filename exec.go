package rnm

type result struct {
	OldPath string
	NewPath string
	Renamed bool
}

func Exec(opts Option) (results []result, err error) {
	targetPaths, err := listFiles(opts.Pattern)

	if err != nil {
		return nil, err
	}

	results = make([]result, len(targetPaths))
	renamer := actualRenamer{}

	for i, path := range targetPaths {
		dirPath, fileName := splitPath(path)

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

		if err != nil {
			return nil, err
		}

		results[i] = result{
			OldPath: oldPath,
			NewPath: newPath,
			Renamed: err == nil,
		}
	}

	return results, nil
}
