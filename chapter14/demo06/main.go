package main

import (
	"fmt"
)

/*
 * @GoName main
 * @Author Crow
 * @Email 648960069@qq.com
 * @Date 2022 10:49
 */

func main() {

	//channel的遍历和关闭
	//1: 使用内置函数close可以关闭channel,当channel关闭后,就不能再向channel写数据 但是可以读数据
	intChan := make(chan int, 3)
	intChan <- 30
	intChan <- 100
	close(intChan) //channel is close
	//intChan <- 20  //panic: send on closed channel
	fmt.Println("ok ok .....")
	//when channel is close , channel can read data
	n1 := <-intChan
	fmt.Println("n1 =", n1)

	//2:使用 range 遍历管道时
	intChan1 := make(chan int, 200)

	for i := 0; i < 200; i++ {
		intChan1 <- i * 2
	}
	//1: 当channel没有关闭的请求下,使用range遍历管道
	//for v := range intChan1 {
	//	fmt.Println("v = ", v)
	//} //all goroutines are asleep - deadlock!  管道中的数据可以正常区别读出  但是会报deallock
	//   原因在于 当管道中没有数据时, range 还想取
	//解决的方案:  关闭channel  底层是在队列的尾部元素打个标签  让range找到这个标签是 就不取了
	close(intChan1)
	for v := range intChan1 {
		fmt.Println(v)
	}
}
