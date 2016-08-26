package rnm

import (
	"path/filepath"
)

func listCandidates(globber Globber, patterns []string) (candidates []string, err error) {
	candidates = []string{}
	for _, pattern := range patterns {
		paths, err := globber.Glob(pattern)
		if err != nil {
			return nil, err
		}
		candidates = append(candidates, paths...)
	}

	candidates = removeDuplicatePaths(candidates)
	return candidates, err
}

func selectTargetPaths(converter converter, candidates []string) []string {
	targets := []string{}

	for _, path := range candidates {
		_, fileName := filepath.Split(path)
		if converter.isTarget(fileName) {
			targets = append(targets, path)
		}
	}

	return targets
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
