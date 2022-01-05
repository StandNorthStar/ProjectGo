package main

/*
有两个不明白点：
	通道怎么触发；
	通道执行顺序；
 */

/*
通道怎么关闭：
	关闭channel的操作原则上应该由发送者完成，因为如果仍然向一个已关闭的channel发送数据，会导致程序抛出panic。而如果由接受者关闭channel，则会遇到这个风险。
	发送者、接收者：个人理解指的是goroutine
*/


/*
channel 通道用于并发goroutine的通信； 因为通道用于goroutine之间的通信，相当于两个 goroutine 通信。
示例0： 常见报错一，在同一个goroutine中执行通道操作。

测试：
在main goroutine定义通道
1. 在main goroutine写入，在goroutine1读出；
2. 在goroutine1 写入，在main goroutine读出；
3. 在goroutine1写入，在goroutine2读出;
*/
/*
func main() {

	//测试1 ： 在main goroutine写入，在goroutine1读出； (这儿执行错误，可以定义有缓冲通道避免错误)
	//ch0 := make(chan int)
	//ch0 <- 15				// main goroutine写入
	//go func() { fmt.Println(<- ch0) }() 	// goroutine1 读出
	//fmt.Println("The End 0-1")
	// 上面方式报错； 读取和写入顺序调整后正常。
	ch0 := make(chan int)
	go func() { fmt.Println( <- ch0 )}()
	ch0 <- 15
	fmt.Println("The End 0-1")
	// 测试1 结束

	// 测试2： 在goroutine1写入， main goroutine读出
	ch0_1 := make(chan int)
	go func() { ch0_1 <- 101 }()
	fmt.Println(<- ch0_1)
	fmt.Println("The End 0-2")
	// 测试2 结束

	//测试3： 在goroutine1写入， goroutine2 读出 。 (虽然没有报错，但是在goroutine2也没有执行)
	ch0_2 := make(chan int)
	go func() { ch0_2 <- 102 }()
	go func() { fmt.Println(<- ch0_2) }()  // 问题：这儿没有读出通道结果。
	fmt.Println("The End 0-3")
	// 测试3 结束
}
*/


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
/*
func main() {


	//var ch = make(chan int)
	//ch <- 10
	//fmt.Println(<-ch)
	//结果：fatal error: all goroutines are asleep - deadlock!
	//原因：因为ch是一个无缓冲的channel,在执行到ch<-10时阻塞到当前goroutine（也就是main函数所在的goroutine），后面打印语句根本没有机会执行。


	// 修复有两种方式：
	// 1. 既然管道无缓冲，那么添加缓冲即可
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
*/


/*
示例六： 有缓冲区通道
 */

/*
func main() {
	ch6 := make(chan int, 3)

	ch6 <- 10
	ch6 <- 11
	ch6 <- 12
	// 定义三个缓冲区，如果已经写满缓冲区(通道满)，继续往通道写入值会报错。
	//ch6 <- 13

	// 通道其实就是一个地址
	// 当定义一个有缓冲区的通道时，可以允许通道只写入、不读出。
	fmt.Println(ch6)

	fmt.Println(<-ch6)
	fmt.Println(<-ch6)
	fmt.Println(<-ch6)

	// 当缓冲区已经没有值时，继续读取会报错。
	//fmt.Println(<-ch6)
}
*/


/*
示例七： 便利通道
注意：在写入通道goroutine中，最后要close()通道。
 */
/*
func main()  {

	// 测试4： 便利通道
	ch7 := make(chan int)
	//go func() { ch7 <- 104; ch7 <- 105}() // 此种方式在下面循环通道时报错；因为通道有两个值104/105在读取完成后通道已经为空了，此时继续读取就会报错。 所以在写入通道的goroutine最后关闭通道。
	go func() { ch7 <- 104; ch7 <- 105;  close(ch7)}()

	for r := range ch7 {
		fmt.Println(r)
	}
	fmt.Println("The End 7")
	// 测试4结束


	// 测试5：
	ch7_1 := make(chan int)

	go func() { ch7_1 <- 106; ch7_1 <- 107;  close(ch7_1) }()
	for  {
		res1, ok := <- ch7_1
		if ok == false {
			fmt.Println("ch7_1 closed", ok)
			break
		}
		fmt.Println(res1, ok)

	}
	fmt.Println("The End ch7_1")
	// 测试5 结束

}
*/


/*
示例八：
	1. 一个goroutine循环接收
	2. 多个goroutine循环接收通道数据；
	3. 多个goroutine同时接收通道的一个数据
 */

func main() {
	//1.  一个goroutine循环接收
	ch8 := make(chan int)
	go func() { ch8 <- 10 }()
}
