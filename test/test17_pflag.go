package main

import (
	"github.com/spf13/pflag"
	"strings"
)

var cliHelp = pflag.StringP("help", "h", "", "Help Messages")
var cliName = pflag.StringP("name", "n", "", "INPUT YOUR NAME")
var cliAge = pflag.IntP("age", "a",1 ,"INPUT YOUR AGE")
var cliDes = pflag.StringP("desc", "d", "", "INPUT DESCRIPTION")
var cliT1 = pflag.StringP("t1", "t","", "hiden parms")

func wordSeqNomailze(f *pflag.FlagSet, name string) pflag.NormalizedName {
	from := []string{"-", "_"}
	to := "."
	for _, seq := range from {
		name = strings.Replace(name, seq, to, -1)
	}
	return pflag.NormalizedName(name)
}


func main() {

	pflag.CommandLine.SetNormalizeFunc(wordSeqNomailze)

	pflag.Lookup("age").NoOptDefVal = "25"

	// 把参数标记为即将废除
	pflag.CommandLine.MarkDeprecated("desc", "please use --desc-to instead")

	// 隐藏参数
	pflag.CommandLine.MarkHidden("t1")
	pflag.Parse()



}