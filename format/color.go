package format

import (
	"fmt"
)

// TODO: Ok, Err

func Green(value string) string {
	return fmt.Sprintf("%s%s%s", "\u001b[32m", value, "\u001b[39m")
}

func Red(value string) string {
	return fmt.Sprintf("%s%s%s", "\u001b[31m", value, "\u001b[39m")
}
