package main

import (
	"fmt"
)
//逻辑运算符
func main(){

	//1基本介绍
	//	用于连接多个条件（一般来讲就是关系表达式），最终的结果也是一个bool值
	
	//2 逻辑运算的说明
	// &&  ||  !

	// &&
	var age = 40

	if age > 30 && age < 50 {
		fmt.Println("ok1")
	}
	if age > 30 && age < 40 {
		fmt.Println("ok2")
	}

	//	&&
	if age > 30 || age < 50 {
		fmt.Println("ok3")
	}
	if age > 30 || age < 40 {
		fmt.Println("ok4")
	}

	// !
	if age > 30  {
		fmt.Println("ok5")
	}
	if !(age > 30) {
		fmt.Println("ok6")
	}

	//4 注意事项和细节说明
	//	&&也叫短路与：如果第一个条件为false，则第二个条件不会判断，最终结果为false
	//	||也叫短路或：如果第一个条件为true，则第二个条件不会判断，最终结果为true

	if age < 30 && test() { // age<30  false   test()不执行
		fmt.Println("ok7")
	}

	if age > 30 || test() {//// age<30  ture   test()不执行
		fmt.Println("ok8")
	}

}

func test() bool {
	fmt.Println("test.......")
	return true
}