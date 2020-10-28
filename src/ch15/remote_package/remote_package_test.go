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
