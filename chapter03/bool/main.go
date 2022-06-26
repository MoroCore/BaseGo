package main

import (
	"fmt"
	"unsafe"
)
func main(){

	//1基本介绍
	//	1)布尔类型也叫bool类型，bool类型数据只允许取值true和false
	//	2)bool类型占1个字节。
	//	3)bool类型适于逻辑运算，一般用于程序流程控制

	var b = false
	fmt.Println("b=",b)
	//注意事项
	//	1:bool类型只占1个字节
	fmt.Printf("b的字节数 %d \n",unsafe.Sizeof(b))
	fmt.Printf("b的字节数 %T  \n",b)
}