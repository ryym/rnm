package rnm

import (
	"path/filepath"
	"strings"
)

func listFiles(patterns []string, opts Option) (targetPaths []string, err error) {
	matchedPaths := []string{}
	for _, pattern := range patterns {
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
