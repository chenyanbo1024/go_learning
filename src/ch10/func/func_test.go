// 函数是一等公民
// 1.可以有多个返回值
// 2.所有参数都是值传递：slice、map、channel会有引用的错觉
// 3.函数可以作为变量的值
// 4.函数可以作为参数和返回值

// defer:延时执行函数
// 1.与C#中的try catch finnaly 中的finnaly

package funcTest

import (
	"fmt"
	"math/rand"
	"testing"
	"time"
)

func returnMultiValue() (int, int) {
	return rand.Intn(10), rand.Intn(20)
}

// 函数执行时间封装（较难点，未理解透彻）
func timeSpent(inner func(op int) int) func(op int) int {
	return func(n int) int {
		start := time.Now()
		ret := inner(n)
		fmt.Println("time spent:", time.Since(start).Seconds())
		return ret
	}
}

func slowFun(op int) int {
	time.Sleep(time.Second * 2)
	return op
}

func Sum(ops ...int) int {
	ret := 0
	for _, op := range ops {
		ret += op
	}
	return ret
}

func TestVarParam(t *testing.T) {
	t.Log(Sum(1, 2, 3, 4))
	t.Log(Sum(1, 2, 3, 4, 5))
}

func Clear() {
	fmt.Println("clear resources.")
}

func TestDefer(t *testing.T) {
	defer Clear()
	fmt.Println("Start")
	// panic("err")
}

func TestFn(t *testing.T) {
	// a, b := returnMultiValue()
	// t.Log(a, b)

	tsSF := timeSpent(slowFun)
	t.Log(tsSF(99))
}
