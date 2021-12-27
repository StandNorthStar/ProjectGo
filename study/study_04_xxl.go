package main
//
//import (
//	xxl "github.com/xxl-job/xxl-job-executor-go"
//	"log"
//)
//
//func main() {
//
//	//xxl.AccessToken("")
//	token := ""
//	exec := xxl.NewExecutor(
//		xxl.ServerAddr("http://12.0.216.196:30031/xxl-job-admin"),
//		xxl.AccessToken(token),            //请求令牌(默认为空)
//		xxl.ExecutorIp("127.0.0.1"),    //可自动获取
//		xxl.ExecutorPort("9999"),       //默认9999（非必填）
//		xxl.RegistryKey("golang-jobs"), //执行器名称
//
//	)
//	exec.Init()
//	//设置日志查看handler
//
//	//注册任务handler
//	//exec.RegTask("task.test", task.Test)
//	//exec.RegTask("task.test2", task.Test2)
//	//exec.RegTask("task.panic", task.Panic)
//	log.Fatal(exec.Run())
//}