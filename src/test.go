package main

import (
	"fmt"
	"net/http"
)

func handler(w http.ResponseWriter,r *http.Request){
	//w表示写回给客户端数据
	//r从客户端(浏览器)读到的数据

	w.Write([]byte("hello world"))
	fmt.Println("URL",r.URL)
	fmt.Println("Method",r.Method)
	fmt.Println("HOST",r.Host)

}

func main(){
	//注册回调函数，该回调函数会在服务器被访问时，自动被调用

	//参数1：客户端访问什么路径返回什么数据
	//参数2：回调函数名-->函数必须是w http.ResponseWriter,r *http.Request
	http.HandleFunc("/",handler)
	//创建web服务器
	http.ListenAndServe("127.0.0.1:8000",nil)
}