package format

import (
	"fmt"
	"github.com/ryym/rnm"
)

func FormatResults(results []rnm.Result, opts *rnm.Option) string {
	successes, fails := separateResults(results)

	styler := new(shellTextStyler)

	condition := formatCondition(opts)
	msgSuccess := formatSuccesses(styler, successes)
	msgFails := formatFails(styler, fails)

	return condition + msgSuccess + msgFails
}

func formatCondition(opts *rnm.Option) string {
	if opts.Dryrun {
		return "(Dryrun)\n"
	}
	return ""
}

func formatSuccesses(styler textStyler, results []rnm.Result) string {
	text := ""
	for _, result := range results {
		message := fmt.Sprintf(
			"%s  %s -> %s",
			styler.Success("\u2713"),
			result.OldPath,
			result.NewPath,
		)
		text += message + "\n"
	}
	return text
}

func formatFails(styler textStyler, results []rnm.Result) string {
	if len(results) == 0 {
		return ""
	}

	text := "\nFAILED:\n"
	for _, result := range results {
		message := fmt.Sprintf(
			"%s  %s -> %s (%s)",
			styler.Error("\u2716"),
			result.OldPath,
			result.NewPath,
			result.Error,
		)
		text += message + "\n"
	}
	return text
}

func separateResults(results []rnm.Result) (successes []rnm.Result, fails []rnm.Result) {
	for _, result := range results {
		if result.Error == nil {
			successes = append(successes, result)
		} else {
			fails = append(fails, result)
		}
	}
	return successes, fails
}
