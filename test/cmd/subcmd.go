package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
)
/*
额外：viper 支持Yaml、Json、 TOML、HCL 等格式
 */

/*
嵌套的子命令
 */

type Flag struct {
	address string
	port string
}
var subFlag = Flag{}

var subcmd = &cobra.Command{
	Use: "subcmd",
	Short: "subcmd test",
	Long: `subcmd test haha 
           subcmd how can do`,
	Args: cobra.OnlyValidArgs,
	//// 如果 subcmd 命令后没有参数，则提示帮助信息
    //Args: func(cmd *cobra.Command, args []string) error {
    //	fmt.Println("args:",args)
	//	if len(args) <= 0 {
	//		cmd.Help()
	//	}
	//	return nil
	//},

	Run: func(cmd *cobra.Command, args []string) {
		//a := viper.GetString("address")
		fmt.Printf("subcmd args %s %s \n", subFlag.address,subFlag.port)

	},
}


func init() {

	// 嵌套子命令； 把被嵌套的子命令subcmd添加到addCmd子命令中。
	addCmd.AddCommand(subcmd)
	//subcmd.Flags().StringSliceVar(&subFlag.address, "address", []string{}, "subcommand address")
	//subcmd.Flags().StringVar(&subFlag.address, "address", "", "subcommand address")
	//subcmd.Flags().StringVar(&subFlag.port, "port", "", "subcommand port" )

	//StringVarP(p *string, name, shorthand string, value string, usage string)
	subcmd.Flags().StringVarP(&subFlag.address, "address", "a", "","subcommand address")
	subcmd.Flags().StringVarP(&subFlag.port, "port", "p", "","subcommand port" )
}


