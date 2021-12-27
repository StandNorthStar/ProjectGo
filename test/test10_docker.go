package main

import (
	"fmt"
	"net"
)

/*
1. 本地执行docker相关命令
2. 远程执行docker相关命令

docker相关执行操作：
1. docker pull
2. docker run
3. docker start
4. docker restart
5. docker inspect
6. docker rm
7. docker push
*/


func main() {
	//type Ano struct {
	//	IPS map[string]net.IP
	//}
	//a := net.IP{'1.1.1.1', '2.2.2.2'}
	//type nn []byte
	//a := nn{'1.2.3.4'}
	//fmt.Println(a)

	a := []byte("haha")
	fmt.Println(a)
	fmt.Println(string(a))

	b := net.IP("1.1.1.1")
	fmt.Println(b)
	fmt.Println(string(b))

}




