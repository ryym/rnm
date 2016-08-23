package format

import (
	"fmt"
)

func Green(value string) string {
	return fmt.Sprintf("%s%s%s", "\u001b[32m", value, "\u001b[39m")
}
