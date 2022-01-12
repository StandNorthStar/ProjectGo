package main

import "fmt"

func main() {

	requests := []int{12,2,3,41,5,6,1}

	for n := range requests {
		go run(n)  // 开启多个协程
	}

	for {
		select{}
	}

}

func run(n int) {


	defer func() {
		if err := recover(); err != nil {
			fmt.Println("ERROR: ", err)
		}
	}()

	// 模拟请求错误
	if n % 5 == 0 {
		panic("request error")
	}
	fmt.Printf("%d\n", n)
}
// https://blog.csdn.net/phpduang/article/details/107989304
// 参考