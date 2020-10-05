// 切片内部结构
// len：元素的个数
// cap：内部数组的容量

// 切片共享存储结构

// 数组 VS 切片
// 容量是否可伸缩：数组（No）、切片（Yes）
// 是否可以进行比较：数组（Yes）、切片（No）

package slice_test

import "testing"

func TestSliceInit(t *testing.T) {
	var s0 []int
	t.Log(len(s0), cap(s0))

	s0 = append(s0, 1)
	t.Log(len(s0), cap(s0))

	s1 := []int{1, 2, 3, 4}
	t.Log(len(s1), cap(s1))

	s2 := make([]int, 3, 5)
	t.Log(len(s2), cap(s2))
	//下面这行代码会有一个错误：index out of range
	// 这是因为我们声明了长度为3,容量为5的数组,所以go只对前三个元素进行初始化
	// t.Log(s2[0], s2[1], s2[2], s2[3])

	s2 = append(s2, 1)
	t.Log(s2[0], s2[1], s2[2], s2[3])
}

// 切片共享存储结构
func TestSliceGrowing(t *testing.T) {
	s := []int{}
	for i := 0; i < 10; i++ {
		s = append(s, i)
		t.Log(len(s), cap(s), s[i])
	}
}

func TestShareSliceMemory(t *testing.T) {
	year := []string{"Jan", "Feb", "Mar", "Apr", "May", "Jun", "Jul", "Aug", "Sep", "Oct", "Nov", "Dec"}
	Q2 := year[3:6]
	t.Log(Q2)

	summer := year[5:8]
	t.Log(summer, len(summer), cap(summer))

	// 因为共享了存储结构，所以修改summer[0]的值，year、Q2的值也会发生变化
	summer[0] = "unknow"
	t.Log(summer, len(summer), cap(summer))
	t.Log(year)
}

func TestSliceCompare(t *testing.T) {
	// a := []int{1, 2, 3, 4}
	// b := []int{1, 2, 3, 4}
	// t.Log(a == b) //error
}
