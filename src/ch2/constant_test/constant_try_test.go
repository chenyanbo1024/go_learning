package constant_test

import (
	"testing"
)

const (
	Monday = 1 + iota
	Tuesday
	Webnesday
)

const (
	Readable = 1 << iota
	Writable
	Excuteable
)

func TestConstantTry(t *testing.T) {
	t.Log(Monday, Tuesday, Webnesday)
}

func TestConstantTry1(t *testing.T) {
	a := 1
	t.Log(a&Readable == Readable, a&Writable == Writable, a&Excuteable == Excuteable)
}
