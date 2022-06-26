package main

import (
	"fmt"
	"time"
)

/*
 * @GoName main
 * @Author Crow
 * @Email 648960069@qq.com
 * @Date 2022 17:01
 */

//第一个协程存储数据到管道里面
func putNum(putchan chan int) {
	for i := 2; i < 1000000; i++ {
		putchan <- i
	}
	close(putchan) //关闭管道
}

//读出是素数管道
func primNum(puuchan chan int, primchan chan int, exitchan chan bool) {
	var flag bool
	for {
		num, ok := <-puuchan //读出数据
		flag = true
		if !ok {
			break
		}
		for i := 2; i < num; i++ {
			if num%i == 0 {
				flag = false
				break
			}
		}
		if flag {
			primchan <- num
		}
	}
	exitchan <- true
}
func main() {
	puchan := make(chan int, 2000)   //需要找的数字
	primchan := make(chan int, 2000) //存储是素数的数字
	exithcan := make(chan bool, 4)   //设计退出的管道
	count := 0
	startTime := time.Now()
	go putNum(puchan)        //写入数据
	for i := 0; i < 4; i++ { //等待处理结果4个协程
		go primNum(puchan, primchan, exithcan) //开启4个协程
	}
	go func() {
		for i := 0; i < 4; i++ {
			<-exithcan
		}
		close(primchan) //关闭素数管道
	}()

	for {
		_, ok := <-primchan
		if !ok {
			break
		}
		count++
	}
	endTime := time.Now()
	fmt.Println("spend time = ", endTime.Sub(startTime))
	fmt.Println("main线程退出")
}
