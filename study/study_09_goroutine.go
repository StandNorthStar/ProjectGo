package main

import "fmt"

/*
示例一：
func main() {
	fmt.Println("start---")

	// 疑问，这两个结果截然不同。
	//ch := make(chan int)
	ch := make(chan int, 5)
	go worker(ch)

	fmt.Println("可以做些什么事情")
	<-ch


	fmt.Println("end ---")

}


func worker(ch chan int) {
	for i:=0;i<=10;i++ {
		fmt.Println(i)
		ch <- i
	}
}
*/

/*
示例二：三个功能不相关的goroutine最后结果要汇总累加到result上
*/
/*
func main() {

	fmt.Println("---start ---")

	ch := make(chan int)

	var result int
	go func() {
		fmt.Println("come into goroutine1")
		var r int
		for i:=1;i<=10;i++ {
			r += i
		}
		fmt.Println("1: r", r)
		ch <- r
	}()

	go func() {
		fmt.Println("come into goroutine2")
		var r int = 1
		for i:=1;i<=10;i++ {
			r *= i
		}
		fmt.Println("2: r", r)
		ch <- r
	}()

	go func() {
		fmt.Println("come into goroutine3")
		ch <- 11
	}()

	for i:=0;i<3;i++ {
		result += <- ch
	}
	fmt.Println("result is :", result)
	fmt.Println("---end---")

}
*/

/*
示例三： 两个goroutine无直接关联，但其中一个先达到某一设定条件便退出或超时退出
*/
/*
func main() {
	fmt.Println("--- start ---")

	ch1 := make(chan int)
	ch2 := make(chan uint64)

	go func() {
		for i:=0;;i++{
			ch1 <- i
		}
		fmt.Println("in goroutine1")
	}()

	go func() {
		var i uint64
		for ;;i++ {
			ch2 <- i
		}
		fmt.Println("in goroutine2:")
	}()

	endCh := false

	for endCh != true {
		select {
			case a:= <- ch1:
				if a > 99 {
					fmt.Println("-- end ch1 --")
					endCh = true
				}
			case b := <- ch2:
				if b == 100 {
					fmt.Println("-- end ch2 --")
					endCh = true
				}
			//case <- time.After(time.Microsecond):
			//	fmt.Println("-- end with timeout --")
			//	endCh = true
		}
	}

	fmt.Println("--- The End ---")
}
*/

/*
示例四：循环100次大概需要1微秒的时间
*/
/*
func main()  {
	fmt.Println("--- start 4---")

	ch := make(chan int)
	go func() {
		for i:=0;i<10;i++ {
			ch <- i
		}
		close(ch)  // 这儿需要关闭channel,否则会引发panic
	}()


	for v := range ch {
		fmt.Println(v)
	}
	fmt.Println("--- end 4 ---")
}
*/

/*
示例五： 引发panic
*/

func main() {

	/*
		var ch = make(chan int)
		ch <- 10
		fmt.Println(<-ch)

		结果：fatal error: all goroutines are asleep - deadlock!
		原因：因为ch是一个无缓冲的channel,在执行到ch<-10时阻塞到当前goroutine（也就是main函数所在的goroutine），后面打印语句根本没有机会执行。
	*/

	/*
		// 修复有两种方式：
	*/
	// 1. 既然管道无缓冲，纳闷添加缓冲即可
	var ch1 = make(chan int, 1)
	ch1 <- 10
	fmt.Println(<-ch1)
	// 2. 管道不要放到同一个goroutine里面。  因为此时ch既有发送也有接收而且不在同一个goroutine里面，此时它们不会相互阻塞
	var ch2 = make(chan int)
	go func() {
		ch2 <- 100
	}()
	a5 := <-ch2
	//fmt.Println(<-ch2)
	fmt.Println(a5)

}
