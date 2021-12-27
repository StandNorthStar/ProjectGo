package main

import (
	"fmt"
	"log"
	"net/http"
	"html"
	"time"
)

type myServer struct {
	Name string
	Age int
}
func (m myServer) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	fmt.Println("okokok") // 展示到后台终端
	fmt.Fprintf(w, "hello hello ,how ary you ", html.EscapeString(r.URL.Path))  // 展示到网页页面中
}

func (m myServer) Haha(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "haha ,how ary you ")
}

// ServeMux 用来为url分配路由
func main() {

	m := myServer{
		"haha1",
		18,
	}
	//路由分配方法一
	http.Handle("/test", m)
	// 路由分配方法二
	http.HandleFunc("/haha", func(w http.ResponseWriter, r *http.Request){
		fmt.Fprintf(w, "hello %q", html.EscapeString(r.URL.Path))
	})
	// 路由分配方法三，注意此种方式需要把 mux添加到http.ListenAndServe(":8000", mux)中。
	mux := http.NewServeMux()
	mux.Handle("/t1", m)
	mux.Handle("/t2", m)
	mux.HandleFunc("/t3", func(w http.ResponseWriter, r *http.Request){
		fmt.Fprintf(w, "he--03 %q", html.EscapeString(r.URL.Path))
	})
	mux.Handle("/t4", http.HandlerFunc(m.Haha))  // 方式四，强制类型转换
	// 注意：路由分配方法1、2 不能和3 同时存在

	// 启动方式一
	//log.Fatal(http.ListenAndServe(":8000", mux))

	// 启动方式二
	s := &http.Server{
		Addr: ":8000",
		Handler: mux,
		ReadTimeout: 10 * time.Second,
		WriteTimeout: 10 * time.Second,
	}
	log.Fatal(s.ListenAndServe())

}