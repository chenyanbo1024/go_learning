// Go接口最佳实践
// 倾向于使用小的接口定义，很多接口只包含一个方法
// 较大的接口定义，可以有多个小接口定义组合而成
// 只依赖与必要功能的最小接口

package emptyInterface

import (
	"fmt"
	"testing"
)

func DoSomething(p interface{}) {
	// 使用switch更合适
	if i, ok := p.(int); ok {
		fmt.Println("Integer", i)
		return
	}
	if i, ok := p.(string); ok {
		fmt.Println("String", i)
		return
	}
	fmt.Println("Unknown Type")
}

func TestEmptyInterfaceAssertion(t *testing.T) {
	DoSomething(10)
	DoSomething("chenyanbo")
	DoSomething(false)
}
