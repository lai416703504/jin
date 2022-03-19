package demo

import (
	"github.com/lai416703504/jin/framework/cobra"
	"log"
)

// BarCommand 代表Bar命令
var BarCommand = &cobra.Command{
	Use:     "bar",
	Short:   "bar的简要说明",
	Long:    "bar的长说明",
	Aliases: []string{"ba", "b"},
	Example: "bar命令的例子",
	RunE: func(c *cobra.Command, args []string) error {
		log.Println("execute bar command")
		return nil
	},
}
