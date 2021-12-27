package main

import (
	"fmt"
	"time"
)
/*获取
 */

func main() {

	const timeformat = "2006-01-02 15:04:05"

	nowTime := time.Now()
	yt := nowTime.AddDate(0,0, -1)
	startTime := time.Date(yt.Year(),yt.Month(),yt.Day(), 0,0,0,0,time.Local)
	endTime := time.Date(yt.Year(),yt.Month(),yt.Day(), 23,59,59,0,time.Local)

	st := startTime.Format(timeformat)
	et := endTime.Format(timeformat)

	fmt.Println(st)
	fmt.Println(et)

}
