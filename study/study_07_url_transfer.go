package main

import (
	"fmt"
	"net/url"
)
/*
url 转义
对于URL中特殊字符，转义为ASCILL对应值
如：http://xxx.0.xxx.xxx:xxx/xxl-job-admin/joblog/pageList?jobGroup=0&jobId=394&logStatus=2&filterTime=2021-12-23 00:00:00 - 2021-12-23 23:59:59
转义后：http://xxx.0.xxx.xxx:xxx/xxl-job-admin/joblog/pageList?jobGroup=0&jobId=394&logStatus=2&filterTime=2021-12-23+00%3A00%3A00+-+2021-12-23+23%3A59%3A59
 */

func main() {
	var urlStr string = "http://xxx.xxx.xxx.xxx:xxx/xxl-job-admin/joblog/pageList?jobGroup=0&jobId=394&logStatus=2&filterTime=2021-12-23 00:00:00 - 2021-12-23 23:59:59"

	newurl, err := url.ParseQuery(urlStr)
	newurl2, err2 := url.Parse(urlStr)
	newurl3, err3 := url.ParseRequestURI(urlStr)

	fmt.Println(newurl, err)
	fmt.Println(newurl2, err2)
	fmt.Println(newurl3, err3)

	fmt.Println("---")
	fmt.Println(newurl2.Query().Encode())

	url01 := "jobGroup=0&jobId=394&logStatus=2&filterTime=2021-12-23 00:00:00 - 2021-12-23 23:59:59"
	s := url.QueryEscape(url01)
	fmt.Println(s)

	// 中意此种方法
	url02 := "2021-12-23 00:00:00 - 2021-12-23 23:59:59"
	s1 := url.QueryEscape(url02)
	fmt.Println(s1)
}