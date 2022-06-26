package main

import (
	"fmt"
	"time"
)

/*
 * @GoName main
 * @Author Crow
 * @Email 648960069@qq.com
 * @Date 2022 22:17
 */
func sysHello() {
	for i := 0; i < 10; i++ {
		time.Sleep(time.Second)
		fmt.Println("Hello,world")
	}
}
func test() {

	defer func() {
		if err := recover(); err != nil {
			fmt.Println("test()发送错误 ", err)
		}
	}()
	var myMap map[int]string
	myMap[0] = "go lang"
}
func main() {

	//捕获协程的panic
	go sysHello()
	go test()

	time.Sleep(time.Second + 5)
}
