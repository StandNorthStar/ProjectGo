package cmd

import (
	"github.com/spf13/cobra"
)

/*
子命令
 */
var addCmd = &cobra.Command{
	Use: "add",
	Short: "add test",
	Long: `add test haha 
               add how can do`,
	Run: func(cmd *cobra.Command, args []string) {
		// 如果 add 命令后没有参数，则提示帮助信息
		if len(args) == 0 {
			cmd.Help()
			return
		}
	},
}

func init() {
	// 把子命令addCmd添加到根命令中。
	rootCmd.AddCommand(addCmd)
}
