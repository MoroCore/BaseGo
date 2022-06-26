package main

import (
	"fmt"
)

/*
 * @GoName main
 * @Author Crow
 * @Email 648960069@qq.com
 * @Date 2022 11:30
 */

func main() {

	intChan := make(chan int, 10000)
	exitChan := make(chan bool, 1)

	go readData(intChan, exitChan)
	go writeData(intChan)
	for {

		_, ok := <-exitChan
		
		if !ok {
			break
		}
	}
}

func writeData(intChan chan int) {
	for i := 1; i <= cap(intChan); i++ {
		intChan <- i
		fmt.Println("write data v = ", i)
	}
	close(intChan)
}
func readData(intChan chan int, exitChan chan bool) {

	for {
		v, ok := <-intChan //ok是close(channel)的结果
		if !ok {
			break
		}
		fmt.Printf("readData 读取数据 =  %v   ok = %T \n ", v, ok)
	}
	exitChan <- true
	close(exitChan)
}
