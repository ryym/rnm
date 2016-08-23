package format

import (
	"fmt"
	"github.com/ryym/rnm"
)

// TODO: Print pretty results.
func FormatResults(results []rnm.Result) string {
	// return "", nil
	// return "\u2318", nil

	text := ""
	for _, result := range results {
		text += formatLine(result) + "\n"
	}

	return text
}

func formatLine(result rnm.Result) string {
	return fmt.Sprintf(
		"%s  %s -> %s",
		Green("\u2713"),
		result.OldPath,
		result.NewPath,
	)
}
