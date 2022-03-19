package command

import (
	"fmt"
	"github.com/lai416703504/jin/framework/cobra"
	"github.com/lai416703504/jin/framework/contract"
	"github.com/lai416703504/jin/framework/util"
)

// initEnvCommand 获取env相关的命令
func initEnvCommand() *cobra.Command {
	envCommand.AddCommand(envListCommand)
	return envCommand
}

// envCommand 获取当前的App环境
var envCommand = &cobra.Command{
	Use:   "env",
	Short: "获取当前的App环境",
	Run: func(c *cobra.Command, args []string) {
		// 获取env环境
		container := c.GetContainer()
		envSerivce := container.MustMake(contract.EnvKey).(contract.Env)
		//打印环境
		fmt.Println("environment:", envSerivce.AppEnv())
	},
}

// envListCommand 获取所有的App环境变量
var envListCommand = &cobra.Command{
	Use:   "list",
	Short: "获取所有的环境变量",
	Run: func(c *cobra.Command, args []string) {
		// 获取env环境
		container := c.GetContainer()
		envService := container.MustMake(contract.EnvKey).(contract.Env)
		envs := envService.All()
		outs := [][]string{}
		for k, v := range envs {
			outs = append(outs, []string{k, v})
		}
		util.PrettyPrint(outs)
	},
}
