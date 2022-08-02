package main

import (
	"fmt"
	"log"
	"net"
	"os"
)

func main() {
	//和服务端进行通信
	conn, err := net.Dial("tcp", "127.0.0.1:8001")

	//判断是否错误，如果有错误那么终止运行
	if err != nil {
		log.Fatal(err)
		return
	}

	//main函数结束的时候关闭建立的连接
	defer conn.Close()

	//从键盘获取用户输入，需要使用stdin不能用Scan因为Scan遇到空格会结束
	go func() {
		str := make([]byte, 4096)
		for {
			n, readErr := os.Stdin.Read(str)
			if readErr != nil {
				log.Fatal(err)
				continue
			}
			//将读取到的数据写给服务器
			conn.Write(str[:n])
		}
	}()

	//读取服务端回发的数据
	for {
		buf := make([]byte, 4096)
		n, Rerr := conn.Read(buf)
		if n == 0 {
			fmt.Println("检测到服务端关闭,客户端关闭连接")
			return
		}
		if Rerr != nil {
			log.Fatal(Rerr)
			return
		}
		fmt.Println("服务端回写数据为:", string(buf[:n]))
	}

	//最后关闭连接
	defer conn.Close()
}
