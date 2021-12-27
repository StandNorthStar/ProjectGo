package main

import (
	"archive/tar"
	"compress/gzip"
	"fmt"
	"io"
	"os"
)
/*
1. 压缩文件
2. 解压文件
 */
func main() {
	/*
	先把/mnt/d/WPSCloud 目录压缩为一个文件 wpscloud.tar.bz2
	再把 wpscloud.tar.bz2 解压到 /tmp/
	 */

	// 解压
	tarfile := "/mnt/d/go1.16.6.linux-amd64.tar.gz"
	ftar, err := os.Open(tarfile)
	if err != nil {
		fmt.Println("openfile error :",tarfile)
	}
	defer ftar.Close()

	br := gzip.NewReader(ftar)
	// //创建空文件，准备写入解压后的数据
	mkd, err := os.Create("/tmp/test01/")
	if err != nil {
		fmt.Println("create dir error:",err.Error())
	}
	defer mkd.Close()
	//写入解压后的数据
	_, err = io.Copy("/tmp", br)
	if err != nil {
		fmt.Println(err.Error())
	}


	// 压缩
	filepath := "/mnt/d/WPSCloud"
	filename := "wpscloud.tar.gz"
	// 创建空文件，准备写入压缩后的数据
	d, _ := os.Create(filename)
	defer d.Close()

	dw := gzip.NewWriter(d)
	defer dw.Close()

	tw := tar.NewWriter(dw)
	defer tw.Close()

	//for _, file := range filepath {
	//
	//}

}



