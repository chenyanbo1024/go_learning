package client

import (
	"ch15/series"
	"testing"
)

func TestPackage(t *testing.T) {
	t.Log(series.GetFibonacciSerie(5))

	// 小写的函数是不能被外访问到
	// t.Log(series.square(5))
}
