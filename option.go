package rnm

type Option struct {
	From     string `short:"f" long:"find" description:"The find string, or regular expression when --regex is set. If not set, the whole filename will be replaced."`
	To       string `short:"r" long:"replace" description:"The replace string. With --regex set, --replace can reference parenthesised substrings from --find with $1, $2, $3 etc. If omitted, defaults to a blank string."`
	AsRegexp bool   `short:"e" long:"regex" description:"When set, --find is intepreted as a regular expression."`
	Dryrun   bool   `short:"d" long:"dry-run" description:"Used for test runs. Set this to do everything but rename the file."`
}

type HelpOption struct {
	Help bool `short:"h" long:"help" description:"Show this help message"`
}

type convertOption struct {
	From string
	To   string
}

type renameOption struct {
	OldPath string
	NewPath string
	Dryrun  bool
}
