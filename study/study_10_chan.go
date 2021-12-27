package main

import "fmt"

/*
channel 通道用于并发goroutine的通信
 */

func t3(c chan int) {
	fmt.Println("t1")

	<-c  // goroutine通过"<-c"来等待t3 goroutine中的“完成事件”
}

func main() {

	// 1.
	c := make(chan int, 2)
	c <- 10
	c <- 11

	fmt.Println( <- c)
	fmt.Println( <- c)

	// 2.
	c1 := make(chan int)
	go func() {
		for i:=1;i<10;i++ {
			c1 <- i

		}
	}()
	fmt.Println(<-c1)

	//for i := range c1 {
	//	fmt.Println(i)
	//}

	// 3.
	/*
		这段代码同样会造成死锁，原因是 channel没有缓冲，相当于channel一直都是满的，所以这里会发生阻塞。下面的goroutine还为创建，所以程序会在此一直阻塞，然后。。。就挂掉了
	*/
	c3 := make(chan int)
	//c3 <- 1003
	//go t3(c3)

	// channel c3上没有任何数据可读的情况下会阻塞等待
	// 3.1 修复
	go t3(c3)

	c3 <- 1003


	// 4.
	c4 := make(chan int)
	for i:=1;i<10;i++ {
		go t4(c4, i)
	}


}

func t4(c chan int, index int) {
	r := <- c

	//<- c
	fmt.Println("This T4 worker:", r)
}