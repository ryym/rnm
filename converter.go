package rnm

import (
	"strings"
)

// TODO: Support regexp.

type converter interface {
	isTarget(fileName string) bool
	convert(fileName string) string
}

type stringConverter struct {
	opts convertOption
}

func (sc stringConverter) isTarget(fileName string) bool {
	return strings.Contains(fileName, sc.opts.From)
}

func (sc stringConverter) convert(fileName string) string {
	return strings.Replace(fileName, sc.opts.From, sc.opts.To, -1)
}
