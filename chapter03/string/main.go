package main

import (
	"fmt"
	"unsafe"
)
func main(){

	//字符串就是一串固定长度的字符连接起来的字符序列。Go的字符串是由单个 字节 连接起来的。Go语言的字符串的字节使用UTF-8编码标识Unicode文本

	//string的基本使用
	var address string = "北京长城"
	fmt.Println(address)
	fmt.Printf("address的字节数 %d \n",unsafe.Sizeof(address)) //16个字节 ？？？
	fmt.Printf("address的的长度 %d \n",len(address)) // 北京长城 长度12   abcd  长度4

	//细节
	// 1)Go语言的字符串的字节使用UTF-8编码标识Unicode文本，这样Golang统一使用UTF-8编码,中文乱码问题不会再困扰程序员。
	// 2)2)字符串一旦赋值了，字符串就不能修改了：在Go中字符串是不可变的。
	var str = "hello"
	//str[0] = 'a' // go中的字符串是不可变的
	//var char string = str[0] //类型不匹配
	char := str[0]
	fmt.Printf("Char的类型为 %T\n", char) //uint8
	fmt.Printf("Char= %c\n", char) //Char= h


	//3)字符串的两种表示形式
	//	(1)双引号,会识别转义字符
	//	(2)反引号，以字符串的原生形式输出，包括换行和特殊字符，可以实现防止攻击、输出源代码等效果
	str2 := "abc\nabc"
	fmt.Println(str2)
	
	str3 := `	//3)字符串的两种表示形式
	//	(1)双引号,会识别转义字符
	//	(2)反引号，以字符串的原生形式输出，包括换行和特殊字符，可以实现防止攻击、输出源代码等效果
	str2 := "abc\nabc"
	fmt.Printf(str2)`
	fmt.Println(str3)

	// 4)字符串拼接方式
	str4 := "Hello" + "World"
	fmt.Println(str4)

	//5)当一行字符串太长时，需要使用到多行字符串，可以如下处理
	str5 := "Hello" + "World" +
			"Hello"
	fmt.Println(str5)
}