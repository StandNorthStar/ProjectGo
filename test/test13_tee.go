package main

import (
	"io"
	"os"
)

/*
使用go实现类型与Linux中 tee的命令
原理：
func TeeReader(r Reader, w Writer) Reader
TeeReader 返回一个 Reader，它将从 r 中读到的数据写入 w 中。所有经由它处理的从 r 的读取都匹配于对应的对 w 的写入。它没有内部缓存，即写入必须在读取完成前完成。任何在写入时遇到的错误都将作为读取错误返回。
 */

/*
扩展：Go中引入的Exception处理：defer, panic, recover
 */

func main() {

	file, err := os.Create("tmp.txt")
	if err != nil {
		panic(err)
	}
	defer file.Close()
	writers := []io.Writer{
		file,
		os.Stdout,
	}
	writer := io.MultiWriter(writers...)
	writer.Write([]byte("Go语言中文网"))

}