package main

import (
	"fmt"
	"log"
	"net"
)

func main() {
	conn, err := net.Dial("tcp", "127.0.0.1:8000")
	if err != nil {
		log.Fatal(err)
		return
	}
	defer conn.Close()

	conn.Write([]byte("Are your read?"))

	buf := make([]byte, 2048)
	n, err := conn.Read(buf)
	if err != nil {
		log.Fatal(err)
		return
	}
	fmt.Println("服务器回发的数据为:", string(buf[:n]))

}
