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
	intChan <- 40
	//intChan <- 98  想管道写数据时,不能超过容量
	var num int
	num = <-intChan
	//查看管道的容量和长度
	fmt.Printf("channel len = %d  cap = %d \n", len(intChan), cap(intChan))
	fmt.Printf("num = %d", num)

	//从管道中读取数据  即从队列中读取第一个元素
	// 细节: 从管道中取数据的前提是  管道中有数据 否则报 deallock
	var num2 int
	num2 = <-intChan
	fmt.Printf("form get channel data ： num2 = %d ", num2)
	// channel使用的注意事项
	//	1: channel中只能存放指定的数据类型
	//  2: channel的数据存放满后，就不能再放了 - deadlock
	//  3: 如果从channel取出数据，可以再放
	//  4: 如果没有使用协程情况下，如果channel数据取完了，再取  回报 deadlock
}
