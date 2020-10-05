// 算术运算符
// Go语言没有前置的 ++、-- (++a)

// 比较运算符（与其他语言基本相同）

// 用 == 比较数组
// 1.相同维数且含有相同个数的元素的数组才可以比较
// 2.每个元素都相同的才相等

// 逻辑运算符

// 位运算符（与其他主要编程语言的差异）
// &^按位置零

// 1 &^ 0  --  1
// 1 &^ 1  --  0
// 0 &^ 1  --  0
// 0 &^ 0  --  0
// 解释：只要右边的操作数为1，则结果肯定为0
//      如果右边的操作数为0，则左边的操作数就是结果

package operate_test

import "testing"

const (
	Readable = 1 << iota
	Writable
	Excuteable
)

func TestBitClear(t *testing.T) {
	a := 7

	a = a &^ Readable
	a = a &^ Excuteable

	t.Log(a&Readable == Readable, a&Writable == Writable, a&Excuteable == Excuteable)
}

// 用 == 比较数组
func TestCompareArray(t *testing.T) {
	a := [...]int{1, 2, 3, 4}
	b := [...]int{1, 3, 4, 5}
	// c := [...]int{1, 2, 3, 4, 5}
	d := [...]int{1, 2, 3, 4}
	e := [...]int{1, 2, 3, 5}
	t.Log(a == b)
	// t.Log(a == c)  //长度不相同的数组做比较，会编译错误
	t.Log(a == d)
	t.Log(a == e)
}
