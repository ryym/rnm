package rnm

import (
	"path/filepath"
	"strings"
)

func listFiles(patterns []string, opts Option) (targetPaths []string, err error) {
	matchedPaths := []string{}
	for _, pattern := range patterns {

		// XXX: Glob returns an empty array when the pattern contains
		// some invalid characters like spaces.
		paths, err := filepath.Glob(pattern)
		if err != nil {
			return nil, err
		}
		matchedPaths = append(matchedPaths, paths...)
	}

	targetPaths = []string{}
	for _, path := range matchedPaths {
		if strings.Contains(path, opts.From) {
			targetPaths = append(targetPaths, path)
		}
	}

	return targetPaths, err
}
