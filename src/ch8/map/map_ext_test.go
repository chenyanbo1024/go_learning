// Map 与 工厂模式
// Map的Value可以是一个方法
// 与 Go 的 Dock Type接口方式一起,可以方便的实现单一方法对象的工厂模式

package mapExt

import "testing"

func TestMapWithFunValue(t *testing.T) {
	m := map[int]func(op int) int{}

	m[1] = func(op int) int { return op }
	m[2] = func(op int) int { return op * op }
	m[3] = func(op int) int { return op * op * op }

	t.Log(m[1](2), m[2](2), m[3](2))
}

// 实现Set
// Go的内置集合中没有Set实现,可以map[type]bool
// 元素的唯一性
// 基本操作：增删改查
func TestMapForSet(t *testing.T) {
	mySet := map[int]bool{}
	n := 1
	mySet[1] = true
	if mySet[n] {
		t.Log("1 is exist")
	} else {
		t.Log("1 is not exist")
	}
}
