package series

import "fmt"

func GetFibonacciSerie(n int) []int {
	ret := []int{1, 1}
	for i := 2; i < n; i++ {
		ret = append(ret, ret[i-2]+ret[i-1])
	}
	return ret
}

// 小写的函数是不能被外访问到
func square(n int) int {
	return n * n
}

// inti方法
// 1.在main被执行前，所有依赖的package的init方法都会被执行
// 2.不同包的init函数按照包导入的依赖关系决定执行顺序
// 3.每个包可以有多个init函数
// 4.包的每个源文件也可以有多个init函数，这点比较特殊

func init() {
	fmt.Println("package init")
}
func init() {
	fmt.Println("package init2")
}
