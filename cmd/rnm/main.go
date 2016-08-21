package main

import (
	"fmt"
	"github.com/jessevdk/go-flags"
	"github.com/ryym/rnm"
	"os"
)

func main() {
	opts := rnm.Option{}
	helpOpts := rnm.HelpOption{}

	parser := makeArgsParser(&opts, &helpOpts)
	args, err := parser.Parse()

	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}

	if len(args) < 1 || helpOpts.Help {
		parser.WriteHelp(os.Stdout)
		return
	}

	// TODO: Accept multiple file names (e.g. rnm -f a -r b dir/*)
	pattern := args[0]
	results, err := rnm.Exec(pattern, opts)

	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}

	fmt.Println(results)
}

func makeArgsParser(opts *rnm.Option, helpOpts *rnm.HelpOption) *flags.Parser {
	parser := flags.NewParser(opts, flags.PassDoubleDash)
	parser.AddGroup("Help Options", helpOpts)
	parser.Usage = "[OPTIONS] <files>"
	return parser
}
