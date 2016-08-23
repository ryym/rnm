package main

import (
	"fmt"
	"github.com/jessevdk/go-flags"
	"github.com/ryym/rnm"
	"github.com/ryym/rnm/format"
	// "os"
)

func main() {
	// opts := rnm.Option{}
	// helpOpts := rnm.HelpOption{}

	// parser := makeArgsParser(&opts, &helpOpts)
	// patterns, err := parser.Parse()

	// if err != nil {
	// 	fmt.Fprintln(os.Stderr, err)
	// 	os.Exit(1)
	// }

	// if len(patterns) < 1 || helpOpts.Help {
	// 	parser.WriteHelp(os.Stdout)
	// 	return
	// }

	// results, err := rnm.Exec(patterns, opts)

	// if err != nil {
	// 	fmt.Fprintln(os.Stderr, err)
	// 	os.Exit(1)
	// }

	results := []rnm.Result{
		{ "OLD", "NEW", nil },
		{ "OLD", "NEW", nil },
		{ "OLD", "NEW", nil },
	}

	fmt.Println(format.FormatResults(results))
}

func makeArgsParser(opts *rnm.Option, helpOpts *rnm.HelpOption) *flags.Parser {
	parser := flags.NewParser(opts, flags.PassDoubleDash)
	parser.AddGroup("Help Options", "", helpOpts)
	parser.Usage = "[OPTIONS] <files>"
	return parser
}
