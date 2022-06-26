package main

import (
	"fmt"
	// "time"
	"errors"
)
// 错误处理
func main() {

	// test1()
	// fmt.Println("main()")
	test2()
}
// 1)在默认情况下，当发生错误后(panic),程序就会退出（崩溃.）
// 2)如果我们希望：当发生错误后，可以捕获到错误，并进行处理，保证程序可以继续执行。还可以在捕获到错误后，给管理员一个提示(邮件,短信。。。）
// 3)这里引出我们要将的错误处理机制
// func test() {

// 	res := 10 / 0

// 	fmt.Println("res=",res)
// }
//1)Go语言追求简洁优雅，所以，Go语言不支持传统的try…catch…finally这种处理。
//2)Go中引入的处理方式为：defer,panic,recover
//3)这几个异常的使用场景可以这么简单描述：Go中可以抛出一个panic的异常，然后在defer中通过recover捕获这个异常，然后正常处理func test() {
func test1() {

	//使用defer + recover 来捕获和处理异常
	defer func() { 
		err := recover() //recover()内置函数，可以捕获到异常说明捕获到错误
		if err != nil {
			fmt.Println("err=",err)
		}
	}()
	num1 := 10
	num2 := 0
	res := num1 / num2
	fmt.Println("res=",res)
}
//自定义错误
//	Go程序中，也支持自定义错误，使用errors.New和panic内置函数。
//	1)errors.New("错误说明"),会返回一个error类型的值，表示一个错误
//  2)panic内置函数,接收一个interface{}类型的值（也就是任何值了）作为参数。可以接收error类型的变量，输出错误信息，并退出程序.

func readConf(name string) (err error) {

	if name == "config.ni" {
		return nil
	}else {
		return errors.New("读取文件错误")
	}
}
func test2(){

	err := readConf("config.ini")
	if err != nil {
		panic(err)
	}

	fmt.Println("test2()....")
}