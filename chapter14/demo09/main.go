package main

import (
	"fmt"
	"time"
)

/*
 * @GoName main
 * @Author Crow
 * @Email 648960069@qq.com
 * @Date 2022 16:44
 */

func main() {

	//channel的阻塞机制
	intChan := make(chan int, 1)
	exitChan := make(chan bool, 1)

	go writeData(intChan)
	go readData(intChan, exitChan)

	for {

		v, ok := <-exitChan
		//查看 在使用协程的情况下 管道的读取请求
		// go的底层 如果exitChan没有数据  就会阻塞次线程
		// v 是 管道的值    ok是从管道中读取数据是否成功
		//  当管道关闭时  读取的  ok = false  v = 根据数据类型缺点
		//  当一个channel  只要有读  就不会阻塞  与读取得速度无关
		fmt.Printf(" v = %v  ok = %v \n", v, ok)
		time.Sleep(time.Second)
		if !ok {
			break
		}
	}
}
func writeData(intChan chan int) {
	for i := 1; i <= cap(intChan); i++ {
		intChan <- i
		//fmt.Println("write data v = ", i)
	}
	close(intChan)
}
func readData(intChan chan int, exitChan chan bool) {

	for {
		_, ok := <-intChan //ok是close(channel)的结果
		if !ok {
			break
		}
	}
	exitChan <- true
	time.Sleep(time.Second * 10)
	close(exitChan)
}
