package main

import (
	"fmt"
	"log"
	"net"
)

func main() {
	//和服务端建立连接
	conn, err := net.Dial("udp", "127.0.0.1:9001")
	if err != nil {
		log.Fatal(err)
		return
	}
	buf := make([]byte, 4096)
	_, err = conn.Write([]byte("我是udp协议"))
	if err != nil {
		log.Fatal(err)
		return
	}
	n, err := conn.Write([]byte("测试数据"))
	if err != nil {
		log.Fatal(err)
		return
	}
	fmt.Println("服务端发送数据为:", string(buf[:n]))
}
