package main

import (
	"fmt"
)
func main(){

	//1 函数的基本概念
	//	为完成某一功能的程序指令(语句)的集合,称为函数
	//	在 Go 中,函数分为: 自定义函数、系统函数
	result :=cal(10,20,'+')
	fmt.Println("result=",result)
}
//使用函数解决前面的计算问题
func cal(n1 float64, n2 float64, opterator byte) float64 {

	var res float64

	switch opterator {
		case '+':
			res = n1 + n2
		case '-':
			res = n1 - n2
		case '*':
			res = n1 * n2
		case '/':
			res = n1 / n2
		default:
			fmt.Println("操作符输入错误")
	}
	return res
}
//包的引出
//	包的三大作用
//		区分相同名字的函数、变量等标识符
//		当程序文件很多时,可以很好的管理项目
//		控制函数、变量等访问范围，即作用域