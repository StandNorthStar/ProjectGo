package main

import (
	"fmt"
	"github.com/fatih/color"
)
/*
参考：https://github.com/fatih/color
 */
func  main() {

	red := color.New(color.FgRed).SprintFunc()
	fmt.Printf("test %s ok", red("haoba"))
}