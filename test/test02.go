package main

import "fmt"

/*
匿名函数，顾名思义就是没有名字的函数
匿名函数: 即在需要使用函数时再定义函数，匿名函数没有函数名只有函数体，函数可以作为一种类型被赋值给函数类型的变量，匿名函数也往往以变量方式传递
特性：在需要使用函数时，再定义函数。
作用：1. 匿名函数用作回调函数； 2. 操作封装
 */


/*
匿名函数 用作 回调函数

 */
func callback(f func(int, int) int) int {
	// 提供工作参数？
	return f(10, 20)
}
func add(x int, y int) int {
	return x + y
}



/*
lambda函数(闭包)：相当于函数里面嵌套函数，类似于Python的装饰器
闭包：闭是封闭（函数内部函数），包是包含（该内部函数对外部作用域而非全局作用域的变量的引用）。
闭包指的是：函数内部函数对外部作用域而非全局作用域的引用。
*/

func getSeq() func() int {
	//定义函数getSeq , 返回一个匿名函数”func() int“
	num := 100
	return func() int {
		num += 1
		return num
	}
}



func main() {
	fmt.Println("haha")
	a := func (test string) {
		fmt.Println(test)
	}
	a("test-haha")

	// 回调
	fmt.Println("匿名函数回调测试")
	fmt.Println(add)
	fmt.Println(callback(add))

	// 闭包
	f1 := getSeq()
	fmt.Println("函数闭包测试")
	fmt.Println(f1())
	fmt.Println(f1())

}