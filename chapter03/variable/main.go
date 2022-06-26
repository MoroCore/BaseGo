package main

import "fmt"

func getVal(num1 int, num2 int)(int, int){

	sum := num1 + num2
	sub := num1 - num2

	return sum, sub
}
//全局变量  在go中函数外部定义变量就是全局变量
var k1 = 100
var k2 = 300

var (
	k3 = 400
	k4 = "Marry"
)

func main(){

	//变量都是其程序的基本组成单位
	// sum,sub都是变量
	var1, var2 := getVal(30, 20)
	fmt.Println("var1=", var1, "var2=", var2)

	// 变量的使用步骤
	// 声明变量    变量赋值      使用变量
	var i int //声明变量 
	i = 10    //变量赋值
	fmt.Println("i=", i) //使用变量


	//(1)第一种：指定变量类型，声明后若不赋值，使用默认值
	var j int
	fmt.Println("j=", j)
	//(2)第二种：根据值自行判定变量类型(类型推导)
	var num2 = 10
	fmt.Println("num2=", num2)
	// (3)第三种：省略var,注意:=左侧的变量不应该是已经声明过的，否则会导致编译错误
	num3 := 99
	fmt.Println("num3=", num3)


	// 4)多变量声明
	// 1 golang如何一次性声明多个同类型变量
	var n1, n2, n3 int
	fmt.Println("n1=",n1 ,"n2=",n2,"n3=",n3)
	// 2 golang如何一次性声明多个类型变量
	var v1, v2, v3 = 12, "Tom", 123.0
	fmt.Println("v1=",v1 ,"v2=",v2 ,"v3=",v3)
	// 3 golang如何一次性声明多个类型变量 类型推到
	m1, m2, m3 := 12, "Tom", 123.0
	fmt.Println("m1=",m1 ,"m2=",m2 ,"m3=",m3)



	// 5)该区域的数据值可以在同一类型范围内不断变化(重点)
	var v9 int = 400
	v9 = 40
	v9 = 50
	//v9 = 80.4 //cannot use 80.4 (untyped float constant) as int value in assignment (truncated)
	//细节 80.0 可以转化为int
	fmt.Println("v9=",v9)

	//6)变量在同一个作用域(在一个函数或者在代码块)内不能重名
	//7)变量=变量名+值+数据类型，这一点请大家注意，变量的三要素
	//8)Golang的变量如果没有赋初值，编译器会使用默认值,比如int默认值0  string默认值为空串，      小数默认为0


	//变量的声明，初始化和赋值
	//初始化细节   初始化赋值  类型可以省略

}

