package main

import (
	"fmt"
	"net"
)

/*
 * @GoName main
 * @Author Crow
 * @Email 648960069@qq.com
 * @Date 2022 14:41
 */

func main() {

	fmt.Println("服务器开启监听......")
	listener, err := net.Listen("tcp", "0.0.0.0:8888")
	if err != nil {
		fmt.Println("listen err = ", err)
		return
	}
	defer listener.Close()
	for {
		fmt.Println("等待客户端连接.....")
		conn, err := listener.Accept()
		if err != nil {
			fmt.Println("accept() err = ", err)
		} else {
			fmt.Printf("Accept() suc con = %v 客户端ip = %v \n", conn, conn.RemoteAddr().String())
		}
		go process(conn)

	}
}
func process(connn net.Conn) {

	for {
		buf := make([]byte, 1024)
		fmt.Printf("服务器在等待客户端%s,发送消息 \n", connn.RemoteAddr().String())

		read, err := connn.Read(buf)
		if err != nil {
			fmt.Printf("客户端退出 err = %v \n", err)
			return
		}
		fmt.Print(string(buf[:read]))
	}
}
