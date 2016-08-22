package rnm

import (
	"github.com/mattn/go-zglob"
	"strings"
)

func listFiles(patterns []string, opts Option) (targetPaths []string, err error) {
	matchedPaths := []string{}
	for _, pattern := range patterns {
		paths, err := zglob.Glob(pattern)
		if err != nil {
			return nil, err
		}
		matchedPaths = append(matchedPaths, paths...)
	}

	matchedPaths = removeDuplicatePaths(matchedPaths)

	targetPaths = []string{}
	for _, path := range matchedPaths {
		if strings.Contains(path, opts.From) {
			targetPaths = append(targetPaths, path)
		}
	}

	return targetPaths, err
}

func removeDuplicatePaths(paths []string) []string {
	pathsAdded := make(map[string]bool)
	uniqPaths := []string{}

	for _, path := range paths {
		if !pathsAdded[path] {
			pathsAdded[path] = true
			uniqPaths = append(uniqPaths, path)
		}
	}

	return uniqPaths
}
