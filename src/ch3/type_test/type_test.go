// 基本数据类型

// 类型转换（与其他编程语言的差异）
// 1.Go语言不支持隐式类型转换
// 2.别名和原有类型也不允许进行隐式类型转换

// 类型的预定义值
// 1.math.MaxInt64
// 2.math.MaxFloat64
// 3.math.MaxUint32
// 4.....

// 指针类型（与其他编程语言的差异）
// 不支持指针运算
// string是值类型，其默认的初始值为空字符串，而不是nil

package teyp_test

import "testing"

type MyInt int64

// 隐私类型转换测试
func TestImplicit(t *testing.T) {
	var a int32 = 1
	var b int64
	b = int64(a) // 显式类型转换

	var c MyInt
	c = MyInt(b)

	t.Log(a, b, c)
}

// 指针
func TestPoint(t *testing.T) {
	a := 1
	aPtr := &a

	// aPtr = aPtr + 1  // 不支持指针运算

	t.Log(a, aPtr)
	t.Logf("%T %T", a, aPtr)
}

func TestString(t *testing.T) {
	var s string // 空字符串，非nil
	t.Log("*" + s + "*")
}
