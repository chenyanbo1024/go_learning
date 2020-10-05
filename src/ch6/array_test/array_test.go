package array_test

import "testing"

func TestArrayInit(t *testing.T) {
	var arr [3]int
	arr1 := [4]int{1, 2, 3, 4}
	arr2 := [...]int{5, 6, 7, 8, 9}
	arr[1] = 10
	t.Log(arr)
	t.Log(arr1)
	t.Log(arr2)
}

// 数组遍历
func TestArrayTravel(t *testing.T) {
	arr3 := [...]int{1, 3, 5, 7, 9, 10}
	for i := 0; i < len(arr3); i++ {
		t.Log(arr3[i])
	}

	for index, el := range arr3 {
		t.Log(index, el)
	}

	// _ 下划线代表占位，
	for _, el := range arr3 {
		t.Log(el)
	}
}

// 数组截取   a[开始索引（包含），结束索引（不包含）]
func TestArraySection(t *testing.T) {
	arr4 := [...]int{1, 2, 3, 4, 5, 6}
	arr4_sec := arr4[:3]
	t.Log(arr4_sec)
}
