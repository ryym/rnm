package rnm

import (
	"regexp"
	"strings"
)

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

type regexpConverter struct {
	opts convertOption
	reg  *regexp.Regexp
}

func newRegexpConverter(opts convertOption) (converter regexpConverter, err error) {
	reg, err := regexp.Compile(opts.From)
	return regexpConverter{opts, reg}, err
}

func (rc regexpConverter) isTarget(fileName string) bool {
	return rc.reg.MatchString(fileName)
}

func (rc regexpConverter) convert(fileName string) string {
	return rc.reg.ReplaceAllString(fileName, rc.opts.To)
}
