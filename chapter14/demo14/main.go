package main

import (
	"fmt"
	"time"
)

/*
 * @GoName main
 * @Author Crow
 * @Email 648960069@qq.com
 * @Date 2022 21:34
 */

func main() {

	intChan := make(chan int, 10)
	go func() {
		for i := 1; i < 10; i++ {
			intChan <- i
		}
	}()
	stringChan := make(chan string, 10)
	go func() {
		for i := 1; i < 10; i++ {
			stringChan <- "hello" + fmt.Sprintf("%d", i)
		}
	}()

	for {
		select {
		case v := <-intChan:
			fmt.Printf("从intChan读取的数据%d \n", v)
			time.Sleep(time.Second)
		case v := <-stringChan:
			fmt.Printf("从stringChan取的数据 %v \n", v)
			time.Sleep(time.Second)
		default:
			fmt.Printf("什么都取不到 \n")
			time.Sleep(time.Second)
			return
		}
	}
}
