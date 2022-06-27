package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
)

/*
 * @GoName main
 * @Author Crow
 * @Email 648960069@qq.com
 * @Date 2022 15:00
 */

func main() {

	conn, err := net.Dial("tcp", "0.0.0.0:8888")
	if err != nil {
		fmt.Println("client dial err = ", err)
		return
	}
	reader := bufio.NewReader(os.Stdin)
	readString, err := reader.ReadString('\n')
	if err != nil {
		fmt.Println("readString err = ", err)
	}
	conn.Write([]byte(readString))
	if err != nil {
		fmt.Println("Conn.Write err = ", err)
	}
	fmt.Printf("客户端发送了%d字节的数据,并退出 ", readString)
}
