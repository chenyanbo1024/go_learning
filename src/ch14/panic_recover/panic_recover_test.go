package panicRecoverTest

import (
	"errors"
	"fmt"
	"testing"
)

// panic用于不可以恢复的错误
// panic退出前会执行defer指定的内容

// panic vs os.Exit
// os.Exit退出时不会调用defer指定的函数
// os.Exit退出时不输出当前调用栈信息

// C#或Java中的try(){}catch{},go语言中在defer中可用recover达到错误恢复的效果

// 最常见的“错误恢复”
// defer func() {
//   if err := recover(); err != nil {
//      log.Error("recoverd panic",err)
//   }
// }()
// 上段代码仅仅是记录了错误日志，却没做其他事情，会有以下后果
// 当心!recover成为恶魔
// 形成僵尸服务进程，导致health check失效
// "Let it Crash" 往往是我们恢复不确定性错误的最好方法

func TestPanicVxExit(t *testing.T) {

	// defer func() {
	// 	fmt.Println("Finally!!")
	// }()

	defer func() {
		if err := recover(); err != nil {
			fmt.Println("recover form", err)
		}
	}()

	fmt.Println("Start")

	// os.Exit(-1)
	panic(errors.New("something wrong!"))
}
