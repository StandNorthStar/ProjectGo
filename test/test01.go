// package GoProject
package main	// 声明包
import "fmt"	// 引入包

// 全局变量不能使用 := 简写变量声明。
// 全局变量可以声明不使用，但是局部变量声明后必须使用。
//aa := 100
var aa int = 100

func init() {
	fmt.Println("one ---")
	//fmt.Println(aa)
	// 常量：在程序运行时，不会被修改的量。标识 const
	const a1 = 10
	const a2 = 3
	const a3 = "changliang - 01"
	fmt.Println(a1*a2)
	fmt.Println(a1/a2)
	fmt.Println(a1%a2)
	fmt.Println(a3)

	// 可以被修改的常量，标识：iota
	const (
		a4 = iota
	    a5 = iota
	    a6 = iota
	)
	const (
		a = iota   //0
		b          //1
		c          //2

		d = "ha"   //独立值，iota += 1
		e          //"ha"   iota += 1
		f = 100    //iota +=1
		g          //100  iota +=1
		h = iota   //7,恢复计数
		i          //8
	)
	fmt.Println(a,b,c,d,e,f,g,h,i)
	fmt.Println("iota: --------------")
	fmt.Println(a4, a5, a6)
}

func main() {
	/* 这是我的第一个简单的程序 */
	fmt.Println("Hello, World!")

	fmt.Println("Hello, World!" + "NiHao")

	// 指定变量类型，如果没有初始化，则变量默认为零值
	var t1 bool = false
	var t2 int = 123
	var t3 float32 = 12.0101
	var t4 string = "haha--heihei"
	var t5 int
	fmt.Println(t1)
	fmt.Println(t2, t3, t4, t5)

	/*
		单引号：rune 类型
		双引号：string类型
		反引号：
	*/
	fmt.Println("-----------------")
	// fmt.Println(rune, '----')
	fmt.Println(`===========`)
	var t6 string = ""
	fmt.Println(t6)
	var t7 = 1
	fmt.Println(t7)

	// var t77 int =10
	// 省略 var 声明变量用法 :=
	t77 := 10
	fmt.Println(t77)

	v1, v2, v3 := t6, t7, t77
	fmt.Println(v1, v2, v3)

	v4 := 8
	v5 := v4
	fmt.Println(&v4)
	fmt.Println(&v5)


	/*
	定义变量未赋值时默认值如下：
	bool      -> false
	numbers -> 0
	string    -> ""

	pointers -> nil
	slices -> nil
	maps -> nil
	channels -> nil
	functions -> nil
	interfaces -> nil

	nil是预定义的标识符，代表指针、通道、函数、接口、映射或切片的零值，也就是预定义好的一个变量
	对于Go来说，map，function，channel都是特殊的指针，指向各自特定的实现
	interface并不是一个指针，它的底层实现由两部分组成，一个是类型，一个值，也就是类似于：(Type, Value).只有当类型和值都是nil的时候，才等于nil。
	 */

	//var t8 bool
	//var t9 int
	//var t10 string
	//var t11 *int
	//var t12 []type
	//var t13 map[]
	//

	a := 1
	b := 2
	if ( a == b ) {
		fmt.Println("OK")
	}else if ( a > b) {
		fmt.Println("ERROR")
	}else if ( a < b){
		fmt.Println("HaoBa")
	}

	if c := "cc" ; c != "dd" {
		fmt.Println("c is not NIL")
	}

	// goto 将控制 转移到 标签 的语句
	for i:=0;i<10;i++ {
		if i == 5 {
			goto HAHA  // HAHA 是一个标签，名称可随意定义
		}
		fmt.Println(i)
	}
	// 标签的定义： 以冒号结尾的单词
	HAHA:

	fmt.Println("HAHA---")

	var j int = 10
	for {
		fmt.Println(j)
		if j == 0 {
			break
		}
		j--
	}
	fmt.Println("FOR-02")
	jj := 0
	//var jj int
	for jj <= 5 {
		fmt.Println(jj)
		jj++
	}

	pp1 := ": hello"
	pp2 := 100
	cc1, cc2 := test(pp1, pp2)
	fmt.Println(cc1)
	fmt.Println(cc2)

	// 闭包	闭包是匿名函数，可在动态编程中使用 （可以理解为python中的装饰器）

	haha := 1.001
	fmt.Println("haha : %f", haha)
	// 方法	方法就是一个包含了接收者的函数
	// 例子： func (variable_name variable_data_type) function_name() [return_type]{
	// func (变量名称 变量数据类型) 函数名称() (返回类型)
	var pp3 Circle
	pp3.radius = 5.00
	getmj := pp3.genArea()
	//fmt.Println("Mian Ji: %f", getmj)
	fmt.Println("Mian Ji: %f", getmj)


	// 数组
	// var sz [3] int
	var sz1 = [3]int{1,2,3}
	sz2 := [3]int{4,5,6}
	fmt.Println(sz1)
	fmt.Println(sz2)
	sz3 := [...]int{10,11,12,14,15} // 长度
	fmt.Println(sz3)
	sz4 := [...]string{10:"haha", 2:"heihei",5:"enheng"} // 指定索引位置赋值初始化数组
	sz4[6] = "haha06"
	fmt.Println(sz4)
	fmt.Println(sz4[10])
	fmt.Println(sz4[2])
	fmt.Println(sz4[5])
	fmt.Println(sz4[6])

	i1 := 0
	var sz5 [10] int
	for {
		if i1>=10 {
			break
		}
		sz5[i1] = i1
		i1++
	}
	fmt.Println(sz5)

	var sz6 [5][5] int
	for i10:=0;i10<5;i10++{
		for i11:=0;i11<5;i11++ {
			sz6[i10][i11] = i11
		}
	}
	fmt.Println(sz6)

	var zz1 *int
	zz2 := 3
	zz1 = &zz2
	fmt.Println(zz1)
	fmt.Println(*zz1)
	var zz3 **int
	zz3 = &zz1
	fmt.Println(zz3)
	fmt.Println(**zz3)

	// 可变参数
	test01(1,2,3,4,5)

	//任意类型参数
	fmt.Println("Para: test02")
	test02("haha", 1, "heihei", []int{3,4,5})

	var aa3 float64 = 3.45
	//aa1 := "1"
	aa2 := int(aa3)
	fmt.Println(aa3)
	fmt.Println(aa2)

	/*
	make 被用来分配引用类型的内存： map, slice, channel
	new 被用来分配除了引用类型的所有其他类型的内存： int, string, array等
	两者作用：主要用来创建分配类型内存
	 */

	var bb01 int
	fmt.Println(&bb01)
}

/* 定义结构体 */
type Circle struct {
	radius float64
}
//该 method 属于 Circle 类型对象中的方法
func (aa Circle) genArea() float64 {
	//aa.radius 即为 Circle 类型对象中的属性
	return 3.14 * aa.radius * aa.radius
}

func test(p string, s int) (string, int){
	fmt.Println("I AM test FUNCTION")
	a := "test function"
	b := a + p
	return b,s

}

func test11(p string, s int) (string, int, error){
	fmt.Println("I AM test FUNCTION")
	//a := "test function"
	//b := a + p
	//return b,s
	//return _, _, nil
	return "1", 1 , nil
}
/*
	可变参数类型
*/
func test01(args ...int) {

	for _,arg := range args{
		fmt.Println("ARG: %d \n", arg)
	}
}
/*
任意类型的可变参数

interface{} 类型，空接口，是导致很多混淆的根源。interface{} 类型是没有方法的接口。由于没有 implements 关键字，所以所有类型都至少实现了 0 个方法，
所以 所有类型都实现了空接口。这意味着，如果您编写一个函数以 interface{} 值作为参数，那么您可以为该函数提供任何值。
 */
func test02(v ...interface{}) {
	for _, val := range v {
		fmt.Println(val)
	}
}
