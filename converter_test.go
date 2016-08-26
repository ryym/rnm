package rnm

import "testing"

func TestStringConverter_convert(t *testing.T) {
	type param struct {
		file   string
		expect string
		opts   convertOption
	}

	testCases := map[string]param{
		"replace": {
			"abc",
			"bbc",
			convertOption{"a", "b", false},
		},
		"replace all matches": {
			"ab-cd-ab-cd",
			"ef-cd-ef-cd",
			convertOption{"ab", "ef", false},
		},
		"don't use regexp": {
			"ggo.go",
			"ggo",
			convertOption{".go", "", false},
		},
		"be case sensitive": {
			"aabbabab",
			"AAbbabab",
			convertOption{"aa", "AA", false},
		},
		"do nothing if no match": {
			"abcde",
			"abcde",
			convertOption{"xyz", "NEW", false},
		},
		"don't care validity of file name": {
			"file/name/*&%",
			"dir/name/*&%",
			convertOption{"file", "dir", false},
		},
	}

	for title, p := range testCases {
		converter := stringConverter{p.opts}
		actual := converter.convert(p.file)
		if actual != p.expect {
			t.Errorf(
				"[%v] expected: %v got: %v",
				title, p.expect, actual,
			)
		}
	}
}

func TestStringConverter_isTarget(t *testing.T) {
	type param struct {
		file   string
		from   string
		expect bool
	}

	testCases := map[string]param{
		"return true if file name is match": {
			"abcde",
			"ab",
			true,
		},
		"don't use regexp": {
			"abcde",
			".bc",
			false,
		},
	}

	for title, p := range testCases {
		converter := stringConverter{
			opts: convertOption{From: p.from, AsRegexp: false},
		}
		actual := converter.isTarget(p.file)
		if actual != p.expect {
			t.Errorf(
				"[%v] expected: %v got: %v",
				title, p.expect, actual,
			)
		}
	}
}
