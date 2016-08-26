package rnm

import (
	"errors"
	"strings"
	"testing"
)

type _mockGlobber struct{}

func (_mockGlobber) Glob(pattern string) (files []string, err error) {
	switch pattern {
	case "a":
		return []string{"a.txt", "a-b.txt"}, nil
	case "b":
		return []string{"b.txt", "a-b.txt"}, nil
	case "c":
		return []string{"c.txt", "c-b.txt"}, nil
	}
	return nil, errors.New("Unkown test pattern")
}

func TestListCandidates(t *testing.T) {
	type param struct {
		patterns []string
		expect   []string
	}

	testCases := map[string]param{
		"lists candidates matched with patterns": {
			patterns: []string{"a", "c"},
			expect:   []string{"a.txt", "a-b.txt", "c.txt", "c-b.txt"},
		},
		"lists candidates uniquely": {
			patterns: []string{"a", "b"},
			expect:   []string{"a.txt", "a-b.txt", "b.txt"},
		},
	}

	for title, p := range testCases {
		actual, err := listCandidates(_mockGlobber{}, p.patterns)
		if err != nil {
			t.Fatal(err)
		}
		if !isSameSlices(actual, p.expect) {
			t.Errorf(
				"[%s] - with: %v, expected: %v, got: %v",
				title, p.patterns, p.expect, actual,
			)
		}
	}
}

type _mockConverter struct {
	from string
}

func (m _mockConverter) isTarget(fileName string) bool {
	return strings.Contains(fileName, m.from)
}

func (_mockConverter) convert(fileName string) string {
	return ""
}

func TestSelectTargetPaths(t *testing.T) {
	type param struct {
		paths  []string
		expect []string
	}

	testCases := map[string]param{
		"lists files matched with opts.From": {
			paths:  []string{"abc.txt", "bcd.txt", "cde.txt"},
			expect: []string{"abc.txt", "bcd.txt"},
		},
		"matches only file names, not directories": {
			paths:  []string{"bc/a", "foo/bc/bar/z", "x/y/z/cbcb"},
			expect: []string{"x/y/z/cbcb"},
		},
	}

	converter := _mockConverter{from: "bc"}
	for title, p := range testCases {
		actual := selectTargetPaths(converter, p.paths)
		if !isSameSlices(actual, p.expect) {
			t.Errorf(
				"[%s] - with: %v, expected: %v, got: %v",
				title, p.paths, p.expect, actual,
			)
		}
	}
}
