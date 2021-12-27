package main

import (
	"fmt"
)

func main() {
	fmt.Println("---")

	var a chan int

	a <- 10

	var dd int
	dd = <- a

	fmt.Println(dd)

	//var r chan int
	//
	//n := 30
	//for i:=0;i<n;i++ {
	//	fmt.Println(i)
	//}


}

