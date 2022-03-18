package demo

// Demo 服务的 key
const Key = "jin:demo"

// Demo 服务的接口
type Service interface {
	GetFoo() Foo
}

// Demo 服务接口定义的一个数据结构
type Foo struct {
	Name string
}
