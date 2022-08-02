package main

import (
	"fmt"
	"log"
	"net"
	"time"
)

func main() {
	//创建一个udp地址结构体
	srvAddr, err := net.ResolveUDPAddr("udp", "127.0.0.1:9000")
	if err != nil {
		log.Fatal(err)
		return
	}

	//创建用于通信的socket
	udpConn, err := net.ListenUDP("udp", srvAddr)
	if err != nil {
		log.Fatal(err)
		return
	}

	//最后关闭连接
	defer udpConn.Close()

	//定义用于保存读取客户端值的byte类型的切片
	buf := make([]byte, 4096)

	//读取客户端发送的数据并返回读取的字节数、客户端地址以及error
	fmt.Println("等待客户端数据写入中......")
	n, clietnAddr, err := udpConn.ReadFromUDP(buf)
	if err != nil {
		log.Fatal(err)
		return
	}
	//模拟处理数据
	fmt.Printf("%v客户端发送数据：%s,", clietnAddr, string(buf[:n]))

	//提取系统当前时间
	daytime := time.Now().String()

	//回写数据给server端
	_, err = udpConn.WriteToUDP([]byte(daytime), clietnAddr)
	if err != nil {
		log.Fatal(err)
		return
	}
	udpConn.Close()
}
