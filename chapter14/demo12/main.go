package main

import "fmt"

/*
 * @GoName main
 * @Author Crow
 * @Email 648960069@qq.com
 * @Date 2022 19:31
 */

func main() {

	//channel使用的细节
	//1:channel声明为只读 或者只写
	// 管道在默认的情况下 是双向的
	var intChan chan int //可读  可写
	intChan <- 10
	<-intChan
	// 声明只写channel
	var intOnlyWriteChan chan<- int
	intOnlyWriteChan = make(chan int, 3)
	intOnlyWriteChan <- 2
	//number := <-intOnlyWriteChan  err
	//fmt.Println("intOnlyWriteChan:=", intOnlyWriteChan) //err  deallock

	//声明只读channel
	var intOnlyReadChan <-chan int
	num2 := <-intOnlyReadChan
	// intOnlyReadChan <- 10  error
	fmt.Println("num2:", num2)
}
