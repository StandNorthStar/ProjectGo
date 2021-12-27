package main

import (
	"fmt"
	"time"
)

func main() {

	// switch
	fmt.Println("-----: switch")
	var v_test1 int
	v_test1 = 2
	switch v_test1 {
	case 1:
		fmt.Println("v_test1 : 1")
	case 2:
		fmt.Println("v-test1: 2")
	default:
		fmt.Println("i don't know")
	}
	// range
	sz01 := [3]string{"sz01", "sz02", "sz03"}
	fmt.Println("-----: range")
	for k, v := range sz01 {
		fmt.Printf("type-k: %s, type-v: %s", k, v)
		fmt.Println(k,v)
	}

	// 结构体
	fmt.Println("-----: struct")
	type struct_test struct {
		id int
		name string
		age int
		describe string
	}
	//struct_test{1, "hly", 18, "haha"}
	//struct_test{2, "hely", 28, "haha1"}
	a1 := struct_test{1, "hly", 18, "haha"}


	fmt.Println(a1)
	fmt.Println(a1.id)
	fmt.Println(a1.name)
	fmt.Println(a1.age)

	str := "hello"
	for i, ch := range str {
		fmt.Println(i, ch)
	}
	/*
	切片和数组的区别：
		数组：在go语言中的数组是属于值类型传递，当我们传递一个数组到一个方法中，改变副本的值并不会修改到原本数组的值。所以得到的数组还是原来的样子。
			如果要对数组进行值修改，只能进行指针操作。

		切片(动态数组)：切片的长度是不固定的，并且切片是可以进行扩容追加长度。

		**注意：声明数组时，方括号内写明了数组的长度或使用...自动计算长度，而声明slice时，方括号内没有任何字符。**
	 */

	// 数组
	fmt.Println("-----: shuzhu")
	ss := [3]int{1,2,3}
	fmt.Println(ss)
	fmt.Printf("Type: %T \n", ss)


	// 切片：切片是对数组的抽象。
	fmt.Println("-----: slince")
	s1 := make([]int, 10) // 或者 s1 := []int{1,2,3} 这个也是声明切片
	s2 := ss[1:2] // 直接获取数组下标的数据，也是切片
	fmt.Println(s1)
	fmt.Println(s2)
	s2 = append(s2, 123) // append 只能往后追加
	s2 = append(s2, 12)
	copy(s2, s1) // copy(目标, 源) 把源数组的值 copy到 目标切片中。 注意：长度不一致按照当前目标的长度为准。
	len_s2 := len(s2)
	cap_s2 := cap(s2)
	fmt.Println("---------")
	fmt.Println(s2)
	fmt.Println(len_s2)
	fmt.Println(cap_s2)

	// 声明[]int 类型的 变量 na
    type na []int
	//na = append(na, 15)
	nn := na{1,2,3}
	fmt.Println("------指定类型-------")
	fmt.Println(nn)
	fmt.Println("------指定类型-------")

	// []byte和string区别
	a := []byte("haha")
	fmt.Println(a)
	fmt.Println(string(a))

	// 集合
	/*
	Map 是一种无序的键值对的集合。
	 */
	var m1 map[string]string
	m1 = make(map[string]string) // 通过make为map分配内存
	m1["h1"] = "haha1"
	m1["h2"] = "haha2"
	for k, v := range m1 {
		fmt.Println(k, v)
	}
	fmt.Println(m1)
	fmt.Println(len(m1))
	delete(m1, "h1")
	fmt.Println(m1)
	fmt.Println(len(m1))


	// 接口类型
	/*
	Go 语言提供了另外一种数据类型即接口，它把所有的具有共性的方法定义在一起，任何其他类型只要实现了这些方法就是实现了这个接口。
	下面实现是多态
	 */
	type phone_t1 interface {
		call()
	}
	var phone phone_t1
	phone =new(NokiaPhone)
	phone.call()

	phone = new(IPhone)
	phone.call()

	// size := unsafe.Sizeof(phone{}) 查看内存占用请用

	// 错误


	// 通道

	/*
	并发 go <func-name>
	 */
	fmt.Println("-----并发-----")
	go say("world")
	say("hello")

	/*
	通道  chan,  <- 用于指定通道方向，未指定为双向通信
	 */


}

/*
接口
 */
type NokiaPhone struct {}
func (nokia NokiaPhone) call(){
	fmt.Println("I AM Nokia")
}
type IPhone struct {}
func (iphone IPhone) call() {
	fmt.Println("I AM IPHONE")
}

/*
并发
 */
func say(s string) {
	for i := 0; i < 5; i++ {
		time.Sleep(100 * time.Millisecond)
		fmt.Println(s)
	}
}

