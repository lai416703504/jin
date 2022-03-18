package demo

import (
	"fmt"
	"github.com/lai416703504/jin/framework"
)

// 服务提供方
type DemoServiceProvider struct {
}

// Register 方法是注册初始化服务实例的方法，这里先暂定为 NewDemoService
func (sp *DemoServiceProvider) Register(container framework.Container) framework.NewInstance {
	return NewDemoService
}

// Boot 方法我们这里我们什么逻辑都不执行, 只打印一行日志信息
func (sp DemoServiceProvider) Boot(container framework.Container) error {
	fmt.Println("demo service boot")
	return nil
}

// IsDefer 方法表示是否延迟实例化，我们这里设置为 true，将这个服务的实例化延迟到第一次 make 的时候
func (sp DemoServiceProvider) IsDefer() bool {
	return true
}

// Params 方法表示实例化的参数。我们这里只实例化一个参数：container，表示我们在 NewDemoService 这个函数中，只有一个参数，container
func (sp DemoServiceProvider) Params(container framework.Container) []interface{} {
	return []interface{}{container}
}

// Name 方法直接将服务对应的字符串凭证返回，在这个例子中就是“jin.demo"
func (sp DemoServiceProvider) Name() string {
	return Key
}
