package main

import "fmt"
import "unsafe"
func main(){

	//用于存放整数值的，比如0,-1,2345等等。

	//int8  int16 int32  int64
	//int8  有符号位 1个字节  范围 -2^7 ~ 2^7-1  
	var j int8 = -128 //cannot use 12666 (untyped int constant) as int8 value in variable declaration
	fmt.Println("j=",j)

	//uint8  uint16 uint32  uint64
	//int8   无符号位 1个字节  范围  0~2^8 -1 
	var j1 uint8 = 255 //cannot use 12666 (untyped int constant) as int8 value in variable declaration
	fmt.Println("j1=",j1)


	//int uint rune byte
	//int 有符号位  32为系统 4个字节 64为系统 8个字节   
	//uint 无符号位  32为系统 4个字节 64为系统 8个字节
	//rune   有符号位  int32      表示一个unicode
	// byte  无       uint8       存储字符
	var a int = 88888
	fmt.Println("a=",a)
	var b uint = 1
	fmt.Println("b=",b)
	var c byte = 'a' //c= 97
	fmt.Println("c=",c)  
	// Println  print line %v \n 默认格式  
	// printf   print format  格式化输出
	fmt.Printf( "a的类型 %T\n",a)  //  %T  type 类型
	// 格式化输出整数 %b binary  二进制    %d   十进制 decimalism
	fmt.Printf( "a的字节数 %d\n",unsafe.Sizeof(a)) //unsafe Go程序类型安全的所有操作 


}