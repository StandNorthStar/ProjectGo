package main

import (
	"fmt"
	"golang.org/x/crypto/ssh"
	"log"
	"time"
)
/*
ssh远程服务器，执行命令并返回结果。
 */
// 定义 ssh连接信息 结构体
type sshInfo struct {

	ipaddr string
	username string
	password string
	port int
	timeout time.Duration
}

// 定义ssh执行命令 结构体组合函数 (可以理解为对象编程中的 方法)
func (si *sshInfo) cmd(command string) string{

	conf := &ssh.ClientConfig{
		Timeout: si.timeout,
		User: si.username,
		Auth: []ssh.AuthMethod{ssh.Password(si.password)},
		HostKeyCallback: ssh.InsecureIgnoreHostKey(),  // 不检查hostkey
	}

	//dial 获取ssh client
	addr := fmt.Sprintf("%s:%d", si.ipaddr, si.port) // 这个可以对字符串做拼接
	sshClient, err := ssh.Dial("tcp", addr, conf)
	if err != nil {
		log.Fatal("创建ssh client 失败",err)
	}
	defer sshClient.Close()  // 注意：ssh执行最后关闭连接

	//创建ssh-session
	session, err := sshClient.NewSession()
	if err != nil {
		log.Fatal("创建ssh Session失败", err)
	}
	defer session.Close()

	// 执行远程命令
	/*
	说明：1. 一个session只能执行一个命令，下面使用同一个session执行commnd,结果只会执行第一条命令
	CombinedOutput方法：返回2个值，(结果,err)
	Output方法：返回2个值，(结果，err)
	Run方法： 返回1个值，如果执行成功返回nil，执行失败返回err.

	文档：
	Start 不管命令是否执行成功均返回 nil
	Run 命令执行成功返回 nil，失败返回 err
	Output 输出StdOut
	CombinedOutput 输出 StdOut 和 StdErr
	 */
	//command := "uptime"
	r1, err := session.CombinedOutput(command)
	//fmt.Printf("执行R1结果：%s \n",string(r1))
	if err != nil {
		log.Fatal("执行R1失败：",err)
	}

	//r2, err := session.Output(command)
	//fmt.Printf("执行R2结果：%s \n",string(r2))
	//if err != nil {
	//	log.Fatal("执行R2失败：",err)
	//}
	//
	//r3 := session.Run(command)
	//fmt.Println(r3)

	return string(r1)
}

func main() {
	/*
	思路： 定义结构体、结构体方法；
	结构体存储ssh连接信息
	结构体组合方法实现ssh 执行cmd命令

	在 main函数中 初始化定义的结构体和结构体方法。
	 */
	var timeout_value time.Duration = (60 * 10) * time.Millisecond
	_ssh := sshInfo{
		ipaddr: "xxx.xxx.xxx.xxx",
		username: "xxx",
		password: "xxx",
		port: 22,
		timeout: timeout_value,
	}
	command := "ls /"
	result := _ssh.cmd(command)
	fmt.Printf("执行结果：%s \n", result)


}

