package main

import (
	"bytes"
	"fmt"
	"io"
	"os"
	"strings"
)

func ReadFrom(reader io.Reader, num int) ([]byte, error) {
	p := make([]byte, num)
	n, err := reader.Read(p)
	if n > 0 {
		return p[:n], nil
	}
	return p, err
}

func main() {
	/*
	示例一、从终端输入
	 */
	fmt.Println("reader01: ")
	d1 := os.Stdout
	data, err := ReadFrom(d1, 10)
	fmt.Println("err:",err)
	if err == nil {
		fmt.Println("data1: ", string(data))
	}

	/*
	示例二、从文件读入
	 */

	filename := "/tmp/my.cnf"
	//r1, err := os.ReadFile(filename)
	//if err != nil {
	//	fmt.Println("error:", err)
	//}
	//fmt.Println("r1:",string(r1))

	f1, err := os.Open(filename)
	if err != nil {
		fmt.Println("e2:", err)
	}
	defer f1.Close()
	fmt.Printf("%T \n", f1)
	r1, err := ReadFrom(f1, 10000)
	fmt.Println(string(r1))

	/*
	示例三、 从字符串读入
	 */
	fmt.Printf("%T \n",strings.NewReader("aaaa1"))
	fmt.Printf("%T \n", bytes.NewBuffer([]byte("aaaaa2")))
	fmt.Printf("%T \n", bytes.NewBufferString("aaaa3"))
	fmt.Printf("%T \n", bytes.NewReader([]byte("aaaa4")))
	//a1 := strings.NewReader("aaaa1")
	//a2 := bytes.NewBuffer([]byte("aaaaa2"))
	//a3 := bytes.NewBufferString("aaaa3")
	a4 := bytes.NewReader([]byte("aaaa4"))
	r2, err := ReadFrom(a4, 20)
	if err != nil {
		fmt.Printf("error: %s \n", err)
	}
	fmt.Println("r2: ", string(r2))
}