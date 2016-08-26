package rnm

import (
	"path/filepath"
	"strings"
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

func listFiles(globber Globber, patterns []string, opts Option) (targetPaths []string, err error) {
	matchedPaths := []string{}
	for _, pattern := range patterns {
		paths, err := globber.Glob(pattern)
		if err != nil {
			return nil, err
		}
		matchedPaths = append(matchedPaths, paths...)
	}

	matchedPaths = removeDuplicatePaths(matchedPaths)

	targetPaths = []string{}
	for _, path := range matchedPaths {
		_, fileName := filepath.Split(path)
		if strings.Contains(fileName, opts.From) {
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
