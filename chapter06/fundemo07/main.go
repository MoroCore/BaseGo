package main

import (
	"fmt"
)

//函数的defer
//	在函数中，程序员经常需要创建资源(比如：数据库连接、文件句柄、锁等)，为了在函数执行完毕后，及时的释放资源，Go的设计者提供defer(延时机制)。
func main() {

	sum(10,20)
}

// 细节
//	1)当go执行到一个defer时，不会立即执行defer后的语句，而是将defer后的语句压入到一个栈中[我为了讲课方便，暂时称该栈为defer栈],然后继续执行函数下一个语句。
//	2)当函数执行完毕后，在从defer栈中，依次从栈顶取出语句执行(注：遵守栈先入后出的机制)，
//	3)在defer将语句放入到栈时，也会将相关的值拷贝同时入栈。请看一段代码：
func sum(n1 int, n2 int) int {

	//当执行到defer时， 暂时不执行，会将defer后面的语句压入到独立的栈(defer栈)
	//当函数执行完毕后，再从defer栈， 按照先入后出的方式出栈，执行
	defer fmt.Println("ok1 n1=",n1) // 3 执行
	defer fmt.Println("ok2 n2=",n2) // 2 执行

	result := n1 + n2
	fmt.Println("ok3 result=",result) // 1 执行 
	
	return result   //ok3 result= 30
					//ok2 n2= 20
					//ok1 n1= 10
}
//	defer最主要的价值是在，当函数执行完毕后，可以及时的释放函数创建的资源。
// 		1)在golang编程中的通常做法是，创建资源后，比如(打开了文件，获取了数据库的链接，或者是锁资源)，可以执行deferfile.Close()deferconnect.Close()
// 		2)在defer后，可以继续使用创建资源.
// 		3)当函数完毕后，系统会依次从defer栈中，取出语句，关闭资源.
// 		4)这种机制，非常简洁，程序员不用再为在什么时机关闭资源而烦心。