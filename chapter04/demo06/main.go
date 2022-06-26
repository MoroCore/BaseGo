package main

import (
	"fmt"
)
//进制
func main(){

	//整数，有四种表示方式
	//二进制 十进制 八进制 以数字0开头表示 十六进制 0-9及A-F 以0x或0X开头表示 此处的A-F不区分大小写

	var i int = 5
	fmt.Printf("%b \n",i) // %b 表示为二进制

	var j int = 011  //八进制
	fmt.Println("j=",j) //默认 10进制输出  %o	表示为八进制

	var k int = 0x11 //16进制
	fmt.Println("k=",k) //%x	表示为十六进制，使用a-f %X	表示为十六进制，使用A-F

	//2 进制转换的介绍

	//	3其它进制转十进制
	// 二进制：110001100转成十进制 
	// var num1 = "110001100"
	
	// var num int = 0 * 0
}