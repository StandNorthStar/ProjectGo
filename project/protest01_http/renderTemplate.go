package main

import (
	"html/template"
	"net/http"
)
/*
问题：
1. 如何渲染html？
2. 渲染出的html文件怎么引用到http.hande中。
 */

func RenderTemplate(w http.ResponseWriter, r *http.Request) {
	//var temp *template.Template
	//temp.New()
	////
	//temp := template.New("test01Template.tpl")
	temp, err := template.ParseFiles("./template/test01.tmpl")
	if err != nil {
		panic("Parse Go Tmpl Error")
	}
	name := "testTital01"
	err = temp.Execute(w, name)
	if err != nil {
		panic("Temp Failed")
	}
}
