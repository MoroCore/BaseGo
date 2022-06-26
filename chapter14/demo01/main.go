package main

import (
	"fmt"
	"time"
)

/**
 * @GoName demo01
 * @Author Crow
 * @Email 648960069@qq.com
 * @Date 2022/6/25 16:53
 */
func main() {

	go testGoRouting()

	for i := 0; i < 10; i++ {
		fmt.Println("hello  world")
		time.Sleep(time.Second)
	}
}

func testGoRouting() {

	for i := 0; i < 20; i++ {
		fmt.Println("test hello world")
		time.Sleep(time.Second)
	}
}
