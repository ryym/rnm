package rnm

type Option struct {
	Pattern  string
	From     string
	To       string
	AsRegexp bool
	Dryrun   bool
}

type convertOption struct {
	From     string
	To       string
	AsRegexp bool
}

type renameOption struct {
	OldPath string
	NewPath string
	Dryrun  bool
}
