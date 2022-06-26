package main

import (
	"fmt"
)

// 指针
func main(){

	//1 基本介绍
	//	1)基本数据类型，变量存的就是值，也叫值类型
	//	2)获取变量的地址，用&，比如：varnumint,获取num的地址：&num分析一下基本数据类型在内存的布局.
	//  地址: 计算机的内部分布标识

	//基本数据类型在内存的布局
	var i = 10
	fmt.Println("i的地址为",&i)
	//	3)指针类型，指针变量存的是一个地址，这个地址指向的空间存的才是值
	var prt *int = &i
	//prt 是一个指针变量
	//prt 的类型是 *int
	//prt 的值是 &i
	fmt.Printf("i的地址为%v\n",prt)

	//	4)获取指针类型所指向的值，使用：*，比如：varptr*int,使用*ptr获取ptr指向的值
	fmt.Printf("i的值为%v\n",*prt)

	//2 案例演示
	//	1)写一个程序，获取一个int变量num的地址，并显示到终端
	//	2)将num的地址赋给指针ptr,并通过ptr去修改num的值.
	var num = 100
	var address = &num

	fmt.Printf("num的地址为%v\n",&num)

	*address = 200
	fmt.Printf("num的值为%v\n",num)

	//4 指针的使用细节
	//	1)值类型，都有对应的指针类型，形式为*数据类型，比如int的对应的指针就是*int,float32对应的指针类型就是*float32,依次类推。
	//	2)值类型包括：基本数据类型int系列,float系列,bool,string、数组和结构体struct

	//7值类型和引用类型
	//	1)值类型：基本数据类型int系列,float系列,bool,string、数组和结构体struc
	//	2)引用类型：指针、slice切片、map、管道chan、interface等都是引用类型
	//   	值类型：变量直接存储值，内存通常在栈中分配
	//		引用类型：变量存储的是一个地址，这个地址对应的空间才真正存储数据(值)，内存通常在堆上分配，当没有任何变量引用这个地址时，该地址对应的数据空间就成为一个垃圾，由GC来回收
}