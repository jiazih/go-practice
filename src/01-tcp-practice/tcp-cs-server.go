package main

import (
	"fmt"
	"log"
	"net"
)

func main() {
	//创建一个用于监听的套接字
	listener, err := net.Listen("tcp", "127.0.0.1:8000")
	if err != nil {
		log.Fatal(err)
		return
	}
	defer listener.Close()
	fmt.Println("服务器等待客户端建立连接")

	//阻塞监听客户端请求
	conn, err := listener.Accept()
	if err != nil {
		log.Println(err)
		return
	}
	defer conn.Close()
	fmt.Println("服务器与客户端成功建立连接")

	//读取客户端发送的数据
	buf := make([]byte, 4096)
	n, err := conn.Read(buf)
	if err != nil {
		log.Fatal(err)
		return
	}
	fmt.Println("读取数据为:", string(buf[:n]))
	conn.Write([]byte("ok"))

}
