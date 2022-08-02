package main

import (
	"fmt"
	"log"
	"net"
	"os"
)

func recvFile(conn net.Conn, fileName string) {
	f, err := os.Create("/Users/jiazihan/Desktop/" + fileName)
	fmt.Println("是否执行")
	if err != nil {
		log.Fatal(err)
	}
	buf := make([]byte, 4096)
	for {
		n, _ := conn.Read(buf)
		if n == 0 {
			fmt.Println("文件接收完成")
			return
		}
		f.Write(buf[:n])
	}
	defer f.Close()
}

func main() {
	listener, err := net.Listen("tcp", "127.0.0.1:10000")
	if err != nil {
		log.Fatal(err)
		return
	}

	conn, err := listener.Accept()
	if err != nil {
		log.Fatal(err)
		return
	}

	buf := make([]byte, 4096)
	//获取文件名
	n, err := conn.Read(buf)
	if err != nil {
		log.Fatal(err)
		return
	}
	fileName := string(buf[:n])
	conn.Write([]byte("ok"))
	recvFile(conn, fileName)
	defer conn.Close()
}
