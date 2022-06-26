package main

import "fmt"

/**
 * @GoName demo04
 * @Author Crow
 * @Email 648960069@qq.com
 * @Date 2022/6/25 19:26
 */

func main() {

	var intChan chan int
	intChan = make(chan int, 3)

	fmt.Printf("intChan的类型 %T, intChan的地址%p \n", intChan, &intChan)

	intChan <- 10
	intChan <- 20
	intChan <- 30
	//intChan <- 98  想管道写数据时,不能超过容量
	var num int
	num = <-intChan
	//查看管道的容量和长度
	fmt.Printf("channel len = %d  cap = %d \n", len(intChan), cap(intChan))
	fmt.Printf("num = %d", num)
}
