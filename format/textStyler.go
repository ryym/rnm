package format

import (
	"fmt"
)

type textStyler interface{
	Success(value string) string
	Error(value string) string
}

type shellTextStyler struct{
	Width int
}

func Hello(s *shellTextStyler) int {
	return s.Width * 2
}

func (shellTextStyler) Success(value string) string {
	return fmt.Sprintf("%s%s%s", "\u001b[32m", value, "\u001b[39m")
}

// XXX: 理解できてない！
func (shellTextStyler) Error(value string) string {
	return fmt.Sprintf("%s%s%s", "\u001b[31m", value, "\u001b[39m")
	// return shellTextStyler.format()
}

// //value string, style string
// func (shellTextStyler) format() string {
// 	return "hey!"
// 	// return fmt.Sprintf("%s\u001b[%sm%s", style, value, "\u001b[39m")
// }
