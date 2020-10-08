// 字符串（与其他主流编程语言的差异）
// 1.string是数据类型，不是引用或指针类型
// 2.string是只读的byte slice，len函数可以它所包含的byte数，并不是字符数
// 3.string的byte数组可以存放任何数据

// Unicode UTF-8
// 1.Unicode是一种字符集（code point）
// 2.UTF-8是Unicode的存储实现（转换为字节序列给规则）

package string

import "testing"

func TestString(t *testing.T) {
	var s string
	t.Log(s)

	s = "hello"
	t.Log(len(s))

	s = "\xE4\xB8\xA5"
	t.Log(s, len(s)) //len求出的是其byte数，并不是字符数
}
