package remote

import (
	"testing"

	cm "github.com/easierway/concurrent_map"
)

func TestConcurentMap(t *testing.T) {
	m := cm.CreateConcurrentMap(99)
	m.Set(cm.StrKey("key"), 10)
	t.Log(m.Get(cm.StrKey("key")))
}

// package
// 1.通过 go get 来获取远程依赖
// 2.通过 go get -u 强制从网络更新远程依赖

// 3.注意代码在github上的组织形式，以适应go get
//   3.1. 直接以代码路径开始，不要有src
//   3.2. 示例：https://github.com/easierway/concurrent_map

// Go未解决的依赖问题
// 1.同一环境下，不同项目使用同一包的不同版本
// 2.无法管理对包的特定版本的依赖
