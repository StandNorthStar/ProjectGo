package main

import (
	"bytes"
	"fmt"
	"io"
	"log"
	"os"
	"strings"
)
/*
须实现：
1. INFO/Warning/ERROR日志级别怎么分类
2. 日志格式怎么自定义。
3. 日志输出定义：输出控制台、保存日志文件。
 */

var (
	Trace   *log.Logger // 记录所有日志
	Info    *log.Logger // 重要的信息
	Warning *log.Logger // 需要注意的信息
	Error   *log.Logger // 非常严重的问题
)


func init() {
	/*
	// New 创建一个新的 Logger。out 参数设置日志数据将被写入的目的地
	// 参数 prefix 会在生成的每行日志的最开始出现
	// 参数 flag 定义日志记录包含哪些属性
	func New(out io.Writer, prefix string, flag int) *Logger

	Discard 是一个 io.Writer 接口，调用它的 Write 方法将不做任何事情并且始终成功返回。当某个等级的日志不重要时，使用 Discard 变量可以禁用这个等级的日志。

	// Stdin、Stdout 和 Stderr 是已经打开的文件，分别指向标准输入、标准输出和标准错误的文件描述符
	 */

	file, err := os.OpenFile("error.log", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		log.Fatalln("Failed to open Error log File: ", err)
	}

	Trace = log.New(os.Stdout,   //ioutil.Discard,   // 不显示日志，相当于禁用日志。
		"[TRACE  ]",
		log.Ldate|log.Ltime|log.Lshortfile)

	Info = log.New(os.Stdout,   // 输出到控制台
		"[INFO   ]",
		log.Ldate|log.Ltime|log.Lshortfile)

	Warning = log.New(os.Stdout,   // 输出到控制台
		"[WARNING]",
		log.Ldate|log.Ltime|log.Lshortfile)  // 输出日志样式： prefix | Ldate | Ltime | 日志文件

	Error = log.New(io.MultiWriter(file, os.Stderr),    // 既把日志输出控制台，也把日志输出到 error.log中
		"[ERROR  ]",
		log.Ldate|log.Ltime|log.Lshortfile)
}


func main() {
	// 示例一 ： 把字符串string转换为 []byte 类型
	a := "playing work"
	fmt.Println(a)
	fmt.Println([]byte(a))

	// 示例二 ： 把字符串string转换为 []byte类型
	var buf bytes.Buffer   // 定义一个 bytes.Buffer 变量
	buf.ReadFrom(strings.NewReader(a))  // strings.NewReader() -> io.Reader类型；(操作Buffer对象；把io.Reader对象)

	log.Println(a)
	log.Println(buf.Bytes())
	log.Println(string(buf.Bytes()))

	Trace.Println("hello 01- trace")
	Info.Println("hello 02 - INFO")
	Warning.Println("hello 03 - Warning")
	Error.Println("hello 04 - Error")
}

