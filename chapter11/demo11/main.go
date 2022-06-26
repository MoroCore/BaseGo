package main

import "fmt"

func main() {

	//断言 将接口转换为具体的类型

	var x interface{}
	var b float64 = 1.1
	x = b //空接口可以接收任意类型
	//将接口转换为具体的类型
	y := x.(float64)
	fmt.Printf("y 的类型 %T 值是 = %v", y, y)
	if y, ok := x.(float32); ok {
		fmt.Println("convert success")
		fmt.Printf(" y 的类型是 %T 值是=%v", y, y)
		fmt.Printf("ok 的类型是 %T 值是=%v", ok, ok)
	} else {
		fmt.Println("convert fail")
		fmt.Printf("ok 的类型是 %T 值是=%v", ok, ok)
	}
}
