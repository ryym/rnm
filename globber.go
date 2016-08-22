package rnm

import (
	"github.com/mattn/go-zglob"
)

type Globber interface {
	Glob(pattern string) (names []string, err error)
}

type zGlobber struct{}

func (zGlobber) Glob(pattern string) (names []string, err error) {
	return zglob.Glob(pattern)
}
