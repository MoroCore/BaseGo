package main

import (
	"fmt"
)

//函数的递归调用
//	一个函数在函数体内又调用了本身，我们称为递归调用
// 递归细节
// 1) 执行一个函数时，就创建一个新的受保护的独立空间(新函数栈)
// 2) 函数的局部变量是独立的，不会相互影响
// 3) 递归必须向退出递归的条件逼近，否则就是无限递归，死龟了:)
// 4) 当一个函数执行完毕，或者遇到 return，就会返回，遵守谁调用，就将结果返回给谁，同时当函数执行完毕或者返回时，该函数本身也会被系统销毁

//题 3：猴子吃桃子问题
//有一堆桃子，猴子第一天吃了其中的一半，并再多吃了一个！以后每天猴子都吃其中的一半，然后
//再多吃一个。当到第十天时，想再吃时（还没吃），发现只有 1 个桃子了。问题：最初共多少个桃子


func monkey(number int) int {

	if number == 10 {
		return 1
	}else {
		return ( monkey(number+1 ) + 1 ) * 2
	}
}


//题 2：求函数值
//	f(1)=3; f(n) = 2*f(n-1)+1

func fun(x int) int {

	if x == 1 {
		return 3
	}else {
		return 2 * fun( x - 1) + 1
	}
}


//题 1：斐波那契数 
func fbn(number int) int{

	if number == 1 || number == 2 {
		return 1
	}else {
		return fbn(number - 1) + fbn(number - 2) 
	}
}
func test(n int) {

	if n > 2 {
		n--
		test(n)
	}
	fmt.Println("n=",n)
}
func main(){

	test(10)

	result := fbn(4)
	fmt.Println("result=",result)
	fmt.Println("f(5)=",fun(5))
	fmt.Println("monkey(1)=",monkey(1))
}