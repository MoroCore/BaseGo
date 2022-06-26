package main

import (
	"fmt"
)
//匿名函数

//	Go支持匿名函数，匿名函数就是没有名字的函数，如果我们某个函数只是希望使用一次，可以考虑使用匿名函数，匿名函数也可以实现多次调用。
func main() {

	//匿名函数使用方式1  在定义匿名函数时就直接调用，这种方式匿名函数只能调用一次。
	res1 := func (n1 int, n2 int) int {
		return n1 + n2
	}(10, 20)
	fmt.Println("res1=",res1)
	//匿名函数使用方式2   将匿名函数赋给一个变量(函数变量)，再通过该变量来调用匿名函数
	sub := func (n1 int,n2 int) int {
		return n1 - n2
	}
	res2 := sub(30, 20)
	fmt.Println("res2=",res2)
	
}
//全局匿名函数
//	如果将匿名函数赋给一个全局变量，那么这个匿名函数，就成为一个全局匿名函数，可以在程序有效。
var (
	fun1 = func (n1 int , n2 int) int {
		return n1 * n2
	}
)

