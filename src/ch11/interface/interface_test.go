// Go接口
// 1.接口为非入侵性，实现不依赖于接口定义
// 2.所以接口的定义可以包含在接口使用者包内

package interfaceTest

import "testing"

type Programmer interface {
	WriteHelloWorld() string
}

type GoProgrammer struct {
}

func (gp *GoProgrammer) WriteHelloWorld() string {
	return "fmt.Println(\"Hello World\")"
}

func TestClient(t *testing.T) {
	var p Programmer
	p = new(GoProgrammer)
	t.Log(p.WriteHelloWorld())
}
