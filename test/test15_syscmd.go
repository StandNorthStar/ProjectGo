package main

import (
	"bytes"
	"fmt"
	"io"
	"log"
	"os"
	"os/exec"
)

/*
Start执行不会等待命令完成就，Run会阻塞等待命令完成。

高级用法(派生子进程)：os.StartProcess


*/
func main() {

	// 1. 直接输出Output()  （组合在一起的stdout/stderr输出）
	r := exec.Command("echo", "-n", "haha0-01")
	//r.Start()
	result, err := r.Output()  // Output 会执行Run操作
	if err != nil {
		panic(err)
	}
	fmt.Println("result:", string(result))
	fmt.Println("--------------------------------------------")


	// 2. 输出方法 CombinedOutput （组合在一起的stdout/stderr输出）
	r2 := exec.Command("ls", "-l")
	result2, err := r2.CombinedOutput() // CombinedOutput 会执行Run操作
	if err != nil {
		panic(err)
	}
	fmt.Println("result2:", string(result2))
	fmt.Println("--------------------------------------------")

	// 3. 标准输出和标准错误分别输出
	r3 := exec.Command("echo", "hah sdf")
	var stdout, stderr bytes.Buffer  // 定义两个变量  bytes.Buffer: 缓存器
	r3.Stdout = &stdout   // 把r2.Stdout 值的地址初始化为 stdout的内存地址
	r3.Stderr = &stderr    // 把r2.Stderr 值的地址初始化为 stderr的内存地址
	err3 := r3.Run()     // 由于定义好命令后未运行，所以要执行Run操作。
	if err3 != nil {
		panic(err3)
	}
	fmt.Printf("out3:%s \n", stdout.String())
	fmt.Printf("err3:%s \n", stderr.String())
	fmt.Println("--------------------------------------------")


	// 4. 把执行结果放入到管道中
	r4 := exec.Command("ls", "-l")
	//var stdout3, stderr3 []byte
	stdoutin, err := r4.StdoutPipe()
	stderrin, err := r4.StderrPipe()
	err04 := r4.Run()
	if err04 != nil {
		panic(err04)
	}
	var buf bytes.Buffer
	buf.ReadFrom(stdoutin)

	fmt.Printf("out4:%s \n", buf.Bytes())
	fmt.Printf("err4:%s \n", stderrin)
	fmt.Println("--------------------------------------------")


	// 5.高级用法，派生子进程。 os.StartProcess
	/*
	os.StartProcess:
		它的第一个参数是要运行的进程，(系统命令和二进制可执行文件)
		第二个参数用来传递选项或参数，
		第三个参数是含有系统环境基本信息的结构体
	 */

	env := os.Environ()
	processAttc := &os.ProcAttr{
		Env: env,

	}
	para := []string{"docker", "version"}
	processI, err := os.StartProcess("/usr/bin/docker", para, processAttc)
	if err != nil {
		panic(err)
	}
	log.Println(processI.Pid)
	fmt.Println("--------------------------------------------")


	// 6.管道
	/*
	两种方式：
	1. command中添加bach -c
	2. io.Pipe() 方法
	还是推荐使用第一种方法
	 */
	//6.1
	r6 := exec.Command("bash", "-c", "echo 'haha0-01 t01'|awk '{print $2}'")
	//r.Start()
	r6_result, err := r6.Output()  // Output 会执行Run操作
	if err != nil {
		panic(err)
	}
	fmt.Println("result:", string(r6_result))

	// 6.2
	r61 := exec.Command("echo", "one two three")
	r62 := exec.Command("awk", "-F", " ", "'{print $2}'")

	p1, p2 := io.Pipe()
	defer p1.Close()
	defer p2.Close()

	r61.Stdout = p2  // r61 从管道的p2端写
	r62.Stdin = p1   // r62 从管理的p1端读

	var buffer bytes.Buffer
	r62.Stdout = &buffer

	r61.Start()
	r62.Start()

	r61.Wait()
	p2.Close()
	r62.Wait()

	io.Copy(os.Stdout, &buffer)
	fmt.Println("--------------------------------------------")

}

