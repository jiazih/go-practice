package main

import (
	"fmt"
	"net"
	"strings"
	"time"
)

//用于初始化用户
type Client struct {
	C    chan string
	Name string
	Addr string
}

//定义全局map用于保存用户
var onlineMap map[string]Client

//创建全局channel，用来接收用户上线信息
var message = make(chan string)

func WriteMsgToClient(clnt Client, conn net.Conn) {
	//监听用户自带channel是否有消息
	for msg := range clnt.C {
		conn.Write([]byte(msg + "\n"))
	}
}

func MakeMsg(client Client, msg string) (buf string) {
	buf = "[" + client.Addr + "]" + client.Name + ": " + msg
	return
}

//主要负责处理用户消息读取、发送、改名、下线等功能
func handleConnect(conn net.Conn) {
	defer conn.Close()
	//获取用户地址
	netAddr := conn.RemoteAddr().String()
	//创建用户,默认用户名是ip+端口
	clnt := Client{make(chan string), netAddr, netAddr}

	//将新连接的用户添加至map中，key是ip+port、值是client结构体
	onlineMap[netAddr] = clnt

	//创建一个用于读取用户上线消息的go程
	go WriteMsgToClient(clnt, conn)

	isQuit := make(chan bool)
	hasData := make(chan bool)

	go func() {
		buf := make([]byte, 4096)
		for {
			n, err := conn.Read(buf)
			if n == 0 {
				fmt.Println("数据读取完毕")
				isQuit <- true
				return
			}
			if err != nil {
				fmt.Println("匿名函数中的read读取客户端数据出错")
				return
			}
			msg := string(buf[:n-1])
			if msg == "who" && len(msg) == 3 {
				conn.Write([]byte("online user list:\n"))
				for _, user := range onlineMap {
					userInfo := user.Addr + user.Name + "\n"
					conn.Write([]byte(userInfo))
				}
			}else if len(msg) >=8 && msg[:6] == "rename" {
				newName := strings.Split(msg,"|")[1]
				clnt.Name = newName
				onlineMap[netAddr] = clnt
				conn.Write([]byte("rename successful"))
			} else {
				message <- msg
			}
			hasData <- true
		}
	}()

	//发送用户上线的消息到全局message通道中
	message <- MakeMsg(clnt, "login")

	for {
		select {
		case <- isQuit:
			close(clnt.C)
			delete(onlineMap,clnt.Addr)
			message <- MakeMsg(clnt,"logout")
			return
		case <- time.After(time.Second * 10):
			delete(onlineMap,clnt.Addr)
			message <- MakeMsg(clnt,"logout")
			return
		case <- hasData:
			//该channel什么都不做，如果检测到用户在发数据，那么重制select
		}
	}
}

//Manager 管理go程，主要用于维护用户map和读取全局channel
func Manager() {
	//初始化onlineMap
	onlineMap = make(map[string]Client)

	//监听全局channel是否有数据,有数据存储值msg，无数据阻塞
	for {
		msg := <-message
		//循环监听channel并发送数据给所有用户
		for _, client := range onlineMap {
			client.C <- msg
		}
	}
}

func main() {
	//创建进程套接字
	listener, err := net.Listen("tcp", "127.0.0.1:8765")
	//判断error是否为空
	if err != nil {
		fmt.Println("套接字创建失败", err)
		return
	}
	defer listener.Close()

	go Manager()

	//循环监听客户端请求
	for {
		//创建通信使用的套接字
		conn, err := listener.Accept()
		if err != nil {
			fmt.Println("通信套接字创建失败", err)
			return
		}
		go handleConnect(conn)
	}
}
