package main

import (
	"fmt"
)
//函数使用的注意事项和细节讨论
//	1) 函数的形参列表可以是多个，返回值列表也可以是多个。
//	2) 形参列表和返回值列表的数据类型可以是值类型和引用类型。
//	3) 函数的命名遵循标识符命名规范，首字母不能是数字，首字母大写该函数可以被本包文件和其它包文件使用，类似 public , 首字母小写，只能被本包文件使用，其它包文件不能使用，类似 privat
//  4) 函数中的变量是局部的，函数外不生效
func test() {
	//n1是test函数的局部变量，只能在test函
	// var n1 int = 10
}
//	5) 基本数据类型和数组默认都是值传递的，即进行值拷贝。在函数内修改，不会影响到原来的值。
func test2(n1 int) {
	n1 = n1 + 100
	fmt.Println("test02() n1 = ",n1)
}
// 6) 如果希望函数内的变量能修改函数外的变量(指的是默认以值传递的方式的数据类型)，可以传入变量的地址&，函数内以指针的方式操作变量。从效果上看类似引用 。
func test3(n1 *int) {
	*n1 = *n1 + 10
	fmt.Println("test3() n1 = ",*n1)
}
//	7) Go 函数不支持函数重载
//  8) 在 Go 中，函数也是一种数据类型，可以赋值给一个变量，则该变量就是一个函数类型的变量了。通过该变量可以对函数调用
func getSum(n1 int, n2 int) int {
	return n1 + n2
}
//	9) 函数既然是一种数据类型，因此在 Go 中，函数可以作为形参，并且调用
func myFun(funvar func(int , int) int , num1 int, num2 int) int {
	return funvar(num1, num2)
}
//	10) 为了简化数据类型定义，Go 支持自定义数据类型
//  	基本语法：type 自定义数据类型名 数据类型 // 理解: 相当于一个别名
// 		type myInt int // 这时 myInt 就等价 int 来使用了.
func test4() {
	type myInt int

	var num1 myInt
	var num2 int

	num1 = 40
	num2 = int(num1) //各位，注意这里依然需要显示转换, go认为myInt和int两个类型
	fmt.Println("num1=",num1,"num2=",num2)
}
// 	11) 支持对函数返回值命名
func getSumAndSub(n1 int, n2 int,) (sum int, sub int) {
	sum = n1 + n2
	sub = n1 - n2
	return
}
//	12) 使用 _ 标识符，忽略返回值
//	13) Go 支持可变参数 
//		如果一个函数的形参列表中有可变参数，则可变参数需要放在形参列表最后。
func sum(n1 int , args... int) int {
	sum := n1
	for i := 0 ; i < len(args) ; i++ {
		sum += args[i]
	}
	return sum
}

func main() {
	// num := 20
	// // test2(num)
	// test3(&num)
	// fmt.Println("main() num= ",num)
	a := getSum
	fmt.Printf("a的类型%T,GetSum的类型是%T \n",a,getSum)

	res := a(10, 40 )//等价于 res := getSum(10, 40)
	fmt.Println("res = ", res)

	result := myFun(getSum, 10 , 20)
	fmt.Println("result = ", result)

	sum := sum(10,20,30,40,90)
	fmt.Println("sum = ", sum)
}