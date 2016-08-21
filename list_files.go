package rnm

import (
	"path/filepath"
	"strings"
)

func listFiles(pattern string, opts Option) (paths []string, err error) {
	matchedPaths, err := filepath.Glob(pattern)
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
