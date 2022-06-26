package main

import "fmt"

//内置函数
func main() {

	//1)len：用来求长度，比如string、array、slice、map、channel
	//2)new：用来分配内存，主要用来分配值类型，比如int、float32,struct...返回的是指针举例说明new的使用：

	num1 := 100
	fmt.Printf("num1的类型%T , num1的值=%v, num的地址%v \n",num1, num1, &num1)

	num2 := new(int) //*int
	//num2的类型%T => *int
	//num2的值=地址0xc04204C098 (这个 地址是系统分配)
	//num2的地址%v =地址0xc04206a020 ( 这个地址是系统分配)
	//num2指向的值= 100
	fmt.Printf( "num2的类型%T，num2的值=%v, num2 的地址%v\n num2这个指针,指向的值=%v",num2, num2, &num2, *num2)


	// 3)make：用来分配内存，主要用来分配引用类型，比如channel、map、slice。这个我们后面讲解。
}