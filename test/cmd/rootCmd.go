package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"os"
)

/*
根命令
 */
var rootCmd = &cobra.Command{
	Use: "hly",
	Short: "hly test",
	Long: `hly test haha 
           hly how can do`,

	ValidArgs: []string{"add", "delete"},
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}