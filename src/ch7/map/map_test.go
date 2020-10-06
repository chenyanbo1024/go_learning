// Map声明

// m := map[string]int{"one":1,"two":2}
// m1 :=  map[string]int{}
// m1["one"] = 1
// m2 := make(map[sting]int, 10 /*初始化容量*/)
// 前两种比较常用，第三种是为某个key赋值

// Map元素访问（与其他语言的差异）
// 在访问的key不存在时，仍会返回0值，不能通过返回nil来判断元素是否存在

package map_test

import "testing"

func TestInitMap(t *testing.T) {
	m1 := map[int]int{1: 1, 2: 2, 3: 3}
	t.Log(m1[2])
	t.Logf("len m1:%d", len(m1))

	m2 := map[int]int{}
	m2[4] = 16
	t.Logf("len m2:%d", len(m2))

	m3 := make(map[int]int, 10)
	t.Logf("len m3:%d", len(m3))
}

func TestAccessNotExistingKey(t *testing.T) {
	m1 := map[int]int{}
	t.Log(m1[1])

	m1[2] = 0
	t.Log(m1[2])

	m1[3] = 10
	// 通过以下if来判断key是否存在
	if v, hadExistKey := m1[3]; hadExistKey {
		t.Logf("key 3`s value is %d", v)
	} else {
		t.Log("key 3 is not exist")
	}

}

// Map遍历
func TestTraveMap(t *testing.T) {
	m1 := map[int]int{1: 3, 2: 5, 3: 7}
	for key, value := range m1 {
		t.Log(key, value)
	}
}
