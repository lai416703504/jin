package command

import (
	"github.com/lai416703504/jin/framework/cobra"
)

// AddKernelCommands will add all command/* to root command
func AddKernelCommands(root *cobra.Command) {
	//root.AddCommand(DemoCommand)
	root.AddCommand(envCommand)

	// cron
	root.AddCommand(initCronCommand())

	// app
	root.AddCommand(initAppCommand())
}
