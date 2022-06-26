package main

import (
	"fmt"
)
//inti 函数的注意事项和细节
//	1) 如果一个文件同时包含全局变量定义，init 函数和 main 函数，则执行的流程全局变量定义->init函数->main 函数

// 面试题：案例如果main.go和utils.go都含有变量定义，init函数时，执行的流程又是怎么样的呢？
// 之上而行执行
var age = test()

// 看到全局变量是先被初始化的，我们这里先写函数
func test() int {
	fmt.Println("test()")
	return 90
}

//init函数,通常可以在init函数中完成初始化工作
func init() {
	fmt.Println("init().....")
}
func main(){

	fmt.Println("main().......age=",age)
}