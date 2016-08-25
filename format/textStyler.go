package format

type textStyler interface {
	Success(value string) string
	Error(value string) string
}

type shellTextStyler struct {
	Width int
}

func (s *shellTextStyler) Success(value string) string {
	return s.format(value, "32")
}

func (s *shellTextStyler) Error(value string) string {
	return s.format(value, "31")
}

func (s *shellTextStyler) format(value string, style string) string {
	return "\u001b[" + style + "m" + value + "\u001b[39m"
}
