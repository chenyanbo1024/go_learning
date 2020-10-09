// 函数是一等公民
// 1.可以有多个返回值
// 2.所有参数都是值传递：slice、map、channel会有引用的错觉
// 3.函数可以作为变量的值
// 4.函数可以作为参数和返回值

package funcTest

import (
	"math/rand"
	"testing"
)

func returnMultiValue() (int, int) {
	return rand.Intn(10), rand.Intn(20)
}

func TestFn(t *testing.T) {
	a, b := returnMultiValue()
	t.Log(a, b)
}
