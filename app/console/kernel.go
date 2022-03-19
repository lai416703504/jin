package console

import (
	"github.com/lai416703504/jin/app/console/command/demo"
	"github.com/lai416703504/jin/framework"
	"github.com/lai416703504/jin/framework/cobra"
	"github.com/lai416703504/jin/framework/command"
	"time"
)

// RunCommand 初始化根 Command 并运行
func RunCommand(container framework.Container) error {
	// 根 Command
	var rootCmd = &cobra.Command{
		// 定义根命令的关键字
		Use: "jin",
		// 简短介绍
		Short: "jin 命令",
		// 根命令的详细介绍
		Long: "jin 框架提供的命令行工具，使用这个命令行工具能很方便执行框架自带命令，也能很方便编写业务命令",
		// 根命令的执行函数
		RunE: func(cmd *cobra.Command, args []string) error {
			cmd.InitDefaultHelpFlag()
			return cmd.Help()
		},
		// 不需要出现 cobra 默认的 completion 子命令
		CompletionOptions: cobra.CompletionOptions{DisableDefaultCmd: true},
	}

	// 为根 Command 设置服务容器
	rootCmd.SetContainer(container)
	// 绑定框架的命令
	command.AddKernelCommands(rootCmd)
	// 绑定业务的命令
	AddAppCommand(rootCmd)
	// 执行 RootCommand
	return rootCmd.Execute()
}

func AddAppCommand(rootCmd *cobra.Command) {
	rootCmd.AddCommand(demo.InitFoo())

	// 启动一个分布式任务调度，调度的服务名称为init_func_for_test，每个节点每5s调用一次Foo命令，抢占到了调度任务的节点将抢占锁持续挂载2s才释放
	rootCmd.AddDistributedCronCommand("bar_func_for_test", "*/5 * * * * *", demo.BarCommand, 2*time.Second)
	//// 每秒调用一次Bar命令
	//rootCmd.AddCronCommand("* * * * * *", demo.BarCommand)
}
