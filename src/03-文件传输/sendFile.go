package main

import (
	"fmt"
	"io"
	"log"
	"net"
	"os"
)

func sendFile(conn net.Conn, filePath string) {
	//打开文件
	f, err := os.Open(filePath)
	if err != nil {
		fmt.Println("文件打开失败")
		return
	}
	//最后关闭文件句柄
	defer f.Close()

	//创建用于保存数据的切片
	buf := make([]byte, 4096)
	for {
		//读取数据
		n, err := f.Read(buf)
		if err != nil {
			if err == io.EOF {
				fmt.Println("文件读取完毕")
			} else {
				fmt.Println("文件传输失败")
			}
			return
		}
		_, err = conn.Write(buf[:n])
		if err != nil {
			log.Fatal(err)
		}
	}
}

func main() {
	//获取命令行参数
	list := os.Args

	if len(list) != 2 {
		fmt.Println("格式为sendFile 文件绝对路径")
	}

	//获取文件路径
	filePath := list[1]

	//单独获取文件名
	fileInfo, err := os.Stat(filePath)
	if err != nil {
		fmt.Println("获取文件名称失败")
		return
	}

	fileName := fileInfo.Name()

	//主动发起连接请求
	conn, err := net.Dial("tcp", "127.0.0.1:10000")
	if err != nil {
		fmt.Println("连接服务器失败")
		return
	}

	//发送文件名给服务端
	_, err = conn.Write([]byte(fileName))
	if err != nil {
		log.Fatal(err)
		return
	}

	//等待服务端回发数据
	buf := make([]byte, 16)
	n, err := conn.Read(buf)
	if err != nil {
		log.Fatal(err)
		return
	}
	//判断客户端是否收到文件名
	if "ok" == string(buf[:n]) {
		//借助网络套接字发送数据
		sendFile(conn, filePath)
	}
	defer conn.Close()
}
