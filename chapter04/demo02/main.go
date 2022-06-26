package main

import "fmt"

//关系运算符
func main(){

	//1基本介绍
	//	1)关系运算符的结果都是bool型，也就是要么是true，要么是false
	//	2)关系表达式经常用在if结构的条件中或循环结构的条件中
	// 	==	!=  >  <  <=  >=  6种
	// 演示
	var n1 int = 9
	var n2 int = 8

	fmt.Println(n1 == n2) //false
	fmt.Println(n1 != n2) //true
	fmt.Println(n1 > n2) //true
	fmt.Println(n1 >= n2) //true
	fmt.Println(n1 < n2) //flase 
	fmt.Println(n1 <= n2) //flase
	
	flag := n1 > n2
	fmt.Println("flag=",flag)

	//2 关系运算符的细节说明
	//	1)关系运算符的结果都是bool型，也就是要么是true，要么是false。
	//	2)关系运算符组成的表达式，我们称为关系表达式：a>b
	//	3)比较运算符"=="不能误写成"="!!
	
}