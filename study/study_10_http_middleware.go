package main

import (
	"log"
	"net/http"
	"time"
)

/*
http 中间件： 此例子 HandlerRoot 中输出一个 “hello”，通过中间件添加 “-- my middleware time--”
中间件使用场景：
	1. 当有几十个路由Handler时，添加特定功能。
	2. 剥离非核心业务功能。


 */

/*
http.Handler，http.HandlerFunc和ServeHTTP的关系：
1. http.HandlerFunc实现了http.Handler这个接口
2. Handler实现了func (ResponseWriter, *Request)即ServeHTTP，所以Handler和http.HandlerFunc()有一致的函数签名,可以将该handler()函数进行类型转换，转为http.HandlerFunc

 */


// 1. 定义初始handler
func HandlerRoot(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("hello"))
}

// 2. 为hanlder添加中间件。
func Middlerware(h http.Handler) http.Handler {
	return http.HandlerFunc(func(writer http.ResponseWriter, request *http.Request) {

		timeNow := time.Now().Format("2006-01-02 15:04:05")
		timeNowFormat, _ := time.Parse("2006-01-02 15:04:05", timeNow)


		writer.Write([]byte(timeNowFormat.String() + "\n"))
		h.ServeHTTP(writer, request)

	})

}


func main() {

	mux := http.NewServeMux()
	//mux.HandleFunc("/", HandlerRoot)
	mux.Handle("/", Middlerware(http.HandlerFunc(HandlerRoot)))

	s := &http.Server{
		Addr: ":8080",
		Handler: mux,
		ReadTimeout: 10 * time.Second,
		WriteTimeout: 10 * time.Second,
	}
	log.Fatal(s.ListenAndServe())

}


/*
参考文档：
http://books.studygolang.com/advanced-go-programming-book/ch5-web/ch5-03-middleware.html
 */
