package rnm

import (
	"path/filepath"
	"strings"
)

func listFiles(opts Option) (paths []string, err error) {
	matchedPaths, err := filepath.Glob(opts.Pattern)
	if err != nil {
		return nil, err
	}

	paths = []string{}
	for _, path := range matchedPaths {
		if strings.Contains(path, opts.From) {
			paths = append(paths, path)
		}
	}

	return paths, err
}
