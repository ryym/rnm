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
			convertOption{"a", "b"},
		},
		"replace all matches": {
			"ab-cd-ab-cd",
			"ef-cd-ef-cd",
			convertOption{"ab", "ef"},
		},
		"don't use regexp": {
			"ggo.go",
			"ggo",
			convertOption{".go", ""},
		},
		"be case sensitive": {
			"aabbabab",
			"AAbbabab",
			convertOption{"aa", "AA"},
		},
		"do nothing if no match": {
			"abcde",
			"abcde",
			convertOption{"xyz", "NEW"},
		},
		"don't care validity of file name": {
			"file/name/*&%",
			"dir/name/*&%",
			convertOption{"file", "dir"},
		},
	}

	for title, p := range testCases {
		converter := newStringConverter(p.opts)
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
		converter := newStringConverter(convertOption{From: p.from})
		actual := converter.isTarget(p.file)
		if actual != p.expect {
			t.Errorf(
				"[%v] expected: %v got: %v",
				title, p.expect, actual,
			)
		}
	}
}

func TestRegexpConverter_convert(t *testing.T) {
	type param struct {
		file   string
		expect string
		opts   convertOption
	}

	testCases := map[string]param{
		"replace all matches": {
			"ab-cd-ab-cd",
			"ef-cd-ef-cd",
			convertOption{"ab", "ef"},
		},
		"use regexp": {
			"ggo.go",
			"",
			convertOption{".go", ""},
		},
		"use regexp 2": {
			"some.file",
			"pre-some.file",
			convertOption{"^", "pre-"},
		},
		"do nothing if no match": {
			"abcde",
			"abcde",
			convertOption{"xyz", "NEW"},
		},
		"don't care validity of file name": {
			"file/name/*&%",
			"dir/name/*&%",
			convertOption{"file", "dir"},
		},
	}

	for title, p := range testCases {
		converter, _ := newRegexpConverter(p.opts)
		actual := converter.convert(p.file)
		if actual != p.expect {
			t.Errorf(
				"[%v] expected: %v got: %v",
				title, p.expect, actual,
			)
		}
	}
}

func TestRegexpConverter_isTarget(t *testing.T) {
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
		"use regexp": {
			"abcde",
			"^.bc",
			true,
		},
		"use regexp 2": {
			"abcde",
			".bc$",
			false,
		},
	}

	for title, p := range testCases {
		converter, _ := newRegexpConverter(convertOption{From: p.from})
		actual := converter.isTarget(p.file)
		if actual != p.expect {
			t.Errorf(
				"[%v] expected: %v got: %v",
				title, p.expect, actual,
			)
		}
	}
}
