package contract

const AppKey = "jin:app"

//APP接口定义
type App interface {
	// Version 定义当前版本
	Version() string
	//BaseFolder 定义项目基础地址
	BaseFloder() string
	// ConfigFolder 定义了配置文件的路径
	ConfigFloder() string
	// LogFolder 定义了日志所在路径
	LogFloder() string
	// ProviderFolder 定义业务自己的服务提供者地址
	ProviderFloder() string
	// MiddlewareFolder 定义业务自己定义的中间件
	MiddlewareFloder() string
	// CommandFolder 定义业务定义的命令
	CommandFloder() string
	// RuntimeFolder 定义业务的运行中间态信息
	RuntimeFloder() string
	// TestFolder 存放测试所需要的信息
	TestFloder() string
}
