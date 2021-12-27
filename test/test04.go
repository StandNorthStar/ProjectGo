package main

import (
	"fmt"
	)
/*
学习： 结构体组合函数
参考：https://www.cnblogs.com/oxspirt/p/10817809.html
 */
type Rect struct { width, length float64 }
func main() {

	// 常量
	const (
		a = iota
		b
		c
		c1 = 100
		d = iota
		e
	)

	const (
		str1 = "haha1"
		str2 = "haha2"
		str3 = "haha3"
	)
	fmt.Println("常量a：", a)
	fmt.Println("常量b：", b)
	fmt.Println("常量c：", c)
	fmt.Println("常量str1：", str1)
	fmt.Println("常量str2：", str2)
	fmt.Println("常量str3：", str3)
	fmt.Println("常量d：", d)
	fmt.Println("常量e：", e)

	// 1. 结构体基础
	var rect Rect
	rect.width = 100
	rect.length = 200
	fmt.Println(rect.width * rect.length)

	// 2. 结构体初始化
	var rect1 = Rect{width: 100, length: 200}
	fmt.Println(rect1.width * rect1.length)

	// 3. 结构体值传递
	var rect2 = Rect{width: 100, length: 200}
	fmt.Println(double_area(rect2))

	// 4. 结构体地址传递
	var rect3 = Rect{width: 100, length: 200}
	fmt.Println(double_area1(&rect3))

	// 5. 结构体组合函数
	/*
	上面我们在main函数中计算了矩形的面积，但是我们觉得矩形的面积如果能够作为矩形结构体的“内部函数”提供会更好。
	这样我们就可以直接说这个矩形面积是多少，而不用另外去取宽度和长度去计算。现在我们看看结构体“内部函数”定义方法：

	注意一点就是定义在结构体上面的函数(function)一般叫做方法(method)。
	 */
	var rect4 = Rect{100, 200}
	fmt.Println("Width:", rect4.width, "Length:", rect4.length, "Area:", rect4.area())
}


func double_area(rect Rect) float64 {
	rect.width *= 2
	rect.length *= 2
	return rect.width * rect.length
}

func double_area1(rect *Rect) float64 {
	rect.width *= 2
	rect.length *= 2
	return rect.width * rect.length
}

func (rect Rect) area() float64 {
	return rect.width * rect.length
}


