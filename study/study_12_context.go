package main

import (
	"fmt"
	"context"
	"k8s.io/apimachinery/pkg/util/rand"
	"time"
)

/*
Context可以用于描述哪些事情？
1. 在某段时间做某些事情；
2. 做某些事情需要花费多长时间。

Context示例： 执行某段代码，当执行到某处时，结束代码程序。
示例一： 吃汉堡比赛，奥特曼每秒吃0-5个，计算吃到第10个时所耗费时间。
 */

func chiHanBao(ctx context.Context) <- chan int {
	/*
	ctx 控制程序结束
	return 吃汉堡个数(通过管道返回结果)
	 */
	c := make(chan int)
	// han bao num
	n := 0
	// time
	t := 0
	go func() {
		for {
			select {
				case <- ctx.Done():
					fmt.Printf("耗时：%d秒，吃了:%d 个汉堡\n", t, n)
					return
				case c <- n:
					incr := rand.Intn(5)
					n += incr
					if n >= 10 {
						n= 10
					}
					t++
					fmt.Printf("奥特曼吃了%d个汉堡\n",n )
			}
		}
	}()
	return c
}

func main() {
	ctx, cancel := context.WithCancel(context.Background())
	eatNum := chiHanBao(ctx)
	for n:= range eatNum {
		if n >= 10 {
			cancel()
			break
		}
	}
	fmt.Println("正在统计结果......")
	time.Sleep(1*time.Second)

}

