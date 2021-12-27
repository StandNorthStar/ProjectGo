package main

import (
	"fmt"
	"time"
)


func main() {
	//rd := 123
	//fmt.Println(big.NewInt(int64(rd)))
	//fmt.Println(new(big.Int).SetInt64(0))

	startingTime := time.Now().UTC()
	time.Sleep(5000 * time.Millisecond)
	endingTime := time.Now().UTC()

	var duration time.Duration = endingTime.Sub(startingTime)
	fmt.Println(duration)
	//var durationAsInt64 = int64(duration)
}