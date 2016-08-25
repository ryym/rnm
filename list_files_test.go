package rnm

import (
	"errors"
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
	case "d":
		return []string{"d/a.txt", "d/d.txt", "d/g.txt"}, nil
	}
	return nil, errors.New("Unkown test pattern")
}

func TestListFiles(t *testing.T) {
	type param struct {
		patterns []string
		opts     Option
		expect   []string
	}

	testCases := map[string]param{
		"lists files matched with patterns": {
			patterns: []string{"a", "c"},
			opts:     Option{From: "txt"},
			expect:   []string{"a.txt", "a-b.txt", "c.txt", "c-b.txt"},
		},
		"lists files matched with opts.From": {
			patterns: []string{"a", "c"},
			opts:     Option{From: "a"},
			expect:   []string{"a.txt", "a-b.txt"},
		},
		"matches only file names, not directories": {
			patterns: []string{"d"},
			opts:     Option{From: "d"},
			expect:   []string{"d/d.txt"},
		},
		"lists files uniquely": {
			patterns: []string{"a", "b"},
			opts:     Option{From: "txt"},
			expect:   []string{"a.txt", "a-b.txt", "b.txt"},
		},
	}

	for title, p := range testCases {
		actual, err := listFiles(_mockGlobber{}, p.patterns, p.opts)
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
