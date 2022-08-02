package main

import (
	"fmt"
	"log"
	"net"
)

func HandleConnect(conn net.Conn) {
	defer conn.Close()
	//获取连接客户端的地址
	addr := conn.RemoteAddr()
	fmt.Println(addr, "成功连接 ")
	buf := make([]byte, 4096)

	//循环读取数据
	for {
		fmt.Println("开始读取")
		n, err := conn.Read(buf)
		fmt.Println("读取成功")
		if string(buf[:n]) == "exit\n" || string(buf[:n]) == "exit\r\n" {
			fmt.Println("客户端发送退出命令，退出服务端")
			return
		}
		if n == 0 {
			fmt.Println("检测到客户端已经退出，服务端关闭连接")
			return
		}
		if err != nil {
			log.Fatal(err)
			return
		}
		fmt.Println(string(buf[:n]))
		conn.Write([]byte("Success"))
	}
}

func main() {
	//创建监听套接字
	listener, err := net.Listen("tcp", "127.0.0.1:8001")
	if err != nil {
		log.Fatal(err)
		return
	}

	//关闭套接字
	defer listener.Close()

	for {
		//创建套接字
		fmt.Println("等待客户端传输数据")
		conn, err := listener.Accept()
		if err != nil {
			log.Fatal(err)
			return
		}
		defer conn.Close()
		go HandleConnect(conn)
	}
}
