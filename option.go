package rnm

type Option struct {
	From     string `short:"f" long:"find" description:"The find string."`
	To       string `short:"r" long:"replace" description:"The replace string."`
	AsRegexp bool   `short:"e" long:"regex" description:"When set, --find is intepreted as a regular expression."`
	Dryrun   bool   `short:"d" long:"dry-run" description:"Used for test runs. Set this to do everything but rename the file."`
}

type HelpOption struct {
	Help bool `short:"h" long:"help" description:"Show this help message"`
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
