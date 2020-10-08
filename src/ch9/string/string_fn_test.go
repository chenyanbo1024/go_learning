package string

import (
	"strconv"
	"strings"
	"testing"
)

func TestStringFn(t *testing.T) {
	s := "A,B,C"
	parts := strings.Split(s, ",")
	t.Log(parts)

	t.Log(strings.Join(parts, "--"))
}

func TestConv(t *testing.T) {
	s := strconv.Itoa(10)
	t.Log(s)
}
