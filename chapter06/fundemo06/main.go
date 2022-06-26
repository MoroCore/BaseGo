package main

import "fmt"
import "strings"

//闭包 闭包就是一个函数和与其相关的引用环境组合的一个整体(实体)
func main(){

	f := addUpper()

	fmt.Println(f(1))
	fmt.Println(f(2))

}

//累加器
func addUpper() func(int)int{
	var n int = 10
	return func (x int) int {
		n = n + x
		return n
	}
}
//闭包的最佳实践
//	1)编写一个函数makeSuffix(suffixstring)可以接收一个文件后缀名(比如.jpg)，并返回一个闭包
//	2)调用闭包，可以传入一个文件名，如果该文件名没有指定的后缀(比如.jpg),则返回文件名.jpg,如果已经有.jpg后缀，则返回原文件名。
//	3)要求使用闭包的方式完成
//	4)strings.HasSuffix,该函数可以判断某个字符串是否有指定的后缀。
func makeSuffix(suffix string) func (fileName string) string {

	fileNameSuffix := suffix

	return func(fileName string) string {
		if !strings.HasSuffix(fileName,suffix) {
			return fileName + fileNameSuffix
		}
		return fileName
	}
}