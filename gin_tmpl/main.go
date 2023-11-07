package main

import (
	"fmt"
	"html/template"
	"net/http"
)

// 遇事不决先写注释
func syahello(w http.ResponseWriter, r *http.Request) {
	//2 解析模板
	t, err := template.ParseFiles("./hello.tmpl")
	if err != nil {
		fmt.Println("parse template failed ,err", err)
		return
	}
	//3 渲染模板
	name := "xiaojixu"
	err = t.Execute(w, name)
	if err != nil {
		fmt.Println("render template failed,err", err)
	}
}
func main() {
	http.HandleFunc("/", syahello)
	err := http.ListenAndServe("127.0.0.1:9000", nil)
	if err != nil {
		fmt.Printf("http server start failed, err:%v/n", err)
		return
	}
}
