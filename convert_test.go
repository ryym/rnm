package rnm

import "testing"

func TestConvertWithoutRegexp(t *testing.T) {
	type param struct {
		file   string
		expect string
		opts   convertOption
	}

	opts := func(from string, to string, regex bool) convertOption {
		return convertOption{from, to, regex}
	}

	testCases := map[string]param{
		"replace": {
			"abc",
			"bbc",
			opts("a", "b", false),
		},
		"replace all matches": {
			"ab-cd-ab-cd",
			"ef-cd-ef-cd",
			opts("ab", "ef", false),
		},
		"don't use regexp": {
			"ggo.go",
			"ggo",
			opts(".go", "", false),
		},
		"be case sensitive": {
			"aabbabab",
			"AAbbabab",
			opts("aa", "AA", false),
		},
		"do nothing if no match": {
			"abcde",
			"abcde",
			opts("xyz", "NEW", false),
		},
		"don't care validity of file name": {
			"file/name/*&%",
			"dir/name/*&%",
			opts("file", "dir", false),
		},
	}

	for title, p := range testCases {
		actual := convert(p.file, p.opts)
		if actual != p.expect {
			t.Errorf(
				"[%v] expected: %v got: %v",
				title, p.expect, actual,
			)
		}
	}
}
