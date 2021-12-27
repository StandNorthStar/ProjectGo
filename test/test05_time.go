package main

import (
	"fmt"
	"time"
)
/*
时间处理
 */
type SSHINFO struct {
	IP string
	username string
	password string
	port int

}

func main() {
	/*
	参考： https://blog.csdn.net/wschq/article/details/80114036
	 */
	start1 := time.Now()
	start2 := time.Now()
	start3 := time.Now().Format("2006-01-02 15:04:05")
	fmt.Println(start1)
	fmt.Printf("%s \n", start2.Sub(start1))
	fmt.Printf("Start3 Cur Time Format: %s \n", start3)

	// 字符串转换为时间类型
	const timeFormat = "2006-01-02 15:04:05" // 时间格式,此格式是固定的。你可以理解为其他语言中的"%Y-%m-%d %H:%M:%S"
	mytime := "2021-01-05 08:31:52"
	//t, _ := time.Parse(timeFormat, "Feb 3, 2013 at 7:54pm (PST)")
	t, _ := time.Parse(timeFormat, mytime)
	fmt.Println(t)

	// 转换时区1 (如果layout已带时区时可直接用Parse)
	t1, _ := time.ParseInLocation("2006-01-02 15:04:05", mytime, time.Local)
	fmt.Println(t1)

	// 转换时区2
	loc, _ := time.LoadLocation("Asia/Shanghai")
	fmt.Println(time.Now().In(loc))

	// 获取指定日期的时间戳
	dt, _ := time.Parse("2006-01-02 15:04:05", "2021-08-06 12:24:51")
	fmt.Println(dt.Unix())

	// 格式化当前日期
	start4 := time.Now().In(loc).Format("2006-01-02 15:04:05")
	dt1, _ := time.Parse("2006-01-02 15:04:05", start4)
	fmt.Printf("当前时间格式化dt1:%s \n",dt1)

	// 格式化当前日期，并指定正确时区
	loc1, _ := time.LoadLocation("Asia/Shanghai")
	start5 := time.Now().Format("2006-01-02 15:04:05")
	dt2, _ := time.ParseInLocation("2006-01-02 15:04:05", start5, loc1)
	fmt.Printf("当前时间格式化dt2:%s \n",dt2)

	// time.Duration
	//start6 := time.Duration
	//fmt.Println(start6)
	Test()

}

func Test() {
	var waitFiveHundredMillisections int64 = 500

	startingTime := time.Now().UTC()
	time.Sleep(10 * time.Millisecond)
	endingTime := time.Now().UTC()

	var duration time.Duration = endingTime.Sub(startingTime)
	var durationAsInt64 = int64(duration)

	if durationAsInt64 >= waitFiveHundredMillisections {
		fmt.Printf("Time Elapsed : Wait[%d] Duration[%d]\n", waitFiveHundredMillisections, durationAsInt64)
	} else {
		fmt.Printf("Time DID NOT Elapsed : Wait[%d] Duration[%d]\n", waitFiveHundredMillisections, durationAsInt64)
	}
}