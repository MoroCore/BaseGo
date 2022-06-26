package main

import "fmt"
import "strconv"
//基本数据类型的相互转换
func main(){

	//Golang和java/c不同，Go在不同类型的变量之间赋值时需要显式转换。也就是说Golang中数据类型不能自动转换

	//2 基本语法
	//	表达式T(v)将值v转换为类型T
	//		T:就是数据类型，比如int32，int64，float32等等
	//		v:就是需要转换的变量
	var i int32 = 100
	//  i----->float
	var n1 float32 = float32(i)
	var n2 int8 = int8(i)
	var n3 int64 = int64(i)
	fmt.Printf("i=%v n1=%v n2=%v n3=%v \n",i,n1,n2,n3)

	//4基本数据类型相互转换的注意事项
	//	1)Go中，数据类型的转换可以是从表示范围小-->表示范围大，也可以范围大--->范围小
	//	2)被转换的是变量存储的数据(即值)，变量本身的数据类型并没有变化！
	fmt.Printf("i的类型%T \n",i) //int32

	//	3) 在转换中，比如将int64转成int8【-128---127】，编译时不会报错，只是转换的结果是按溢出处理，和我们希望的结果不一样。因此在转换时，需要考虑范围.
	// var num1 int64 = 99999999
	// var num2 int8 = int8(num1)
	// fmt.Println("num2=",num2)

	//基本介绍
	//在程序开发中，我们经常将基本数据类型转成string,或者将string转成基本数据类型。
	//1基本类型转string类型
	//fmt.Sprintf("%参数",表达式)  Springf 根据format参数生成格式字符串并将字符串返回


	//基本数据类型转成string
	var num1 int = 99
	var num2 float64 = 23.999
	var b bool = true
	var char byte = 'h' //uint8
	var str string
	str = fmt.Sprintf("%d",num1) //%d  表示为十进制
	fmt.Printf("str type %T str = %q \n",str,str)

	str = fmt.Sprintf("%f",num2) //%f 有小数部分但无指数部分，如123.456
	fmt.Printf("str type %T str = %q \n",str,str)

	str = fmt.Sprintf("%t",b) //%t 单词true或false
	fmt.Printf("str type %T str = %q \n",str,str)

	str = fmt.Sprintf("%c",char) //%c  字符
	fmt.Printf("str type %T str = %q \n",str,str)

	//方式2  使用strconv包的函数 strconv包实现了基本数据类型和其字符串表示的相互转换。
	var num3 int = 99
	var num4 float64 = 23.456
	var b2 bool = true
	var num5 uint16 =  100

	str =strconv.FormatInt(int64(num3),10) //func FormatInt(i int64, base int) string  
										   //返回i的base进制的字符串表示。base 必须在2到36之间，结果中会使用小写字母'a'到'z'表示大于10的数字
	fmt.Printf("str type %T str = %q \n",str,str)

	str =strconv.FormatFloat(num4,'f',10,64) //func FormatFloat(f float64, fmt byte, prec, bitSize int) string 
											 //
	fmt.Printf("str type %T str = %q \n",str,str)

	str =strconv.FormatBool(b2) //func FormatBool(b bool) string   根据b的值返回"true"或"false"。
	fmt.Printf("str type %T str = %q \n",str,str)

	str = strconv.FormatUint(uint64(num5),10)
	fmt.Printf("str type %T str = %q \n",str,str)

	//string类型转基本数据类型

	//func ParseBool(str string) (value bool, err error)
	//返回字符串表示的bool值。它接受1、0、t、f、T、F、true、false、True、False、TRUE、FALSE；否则返回错误。
	var str1 string = "true"
	b3, _ := strconv.ParseBool(str1)
	fmt.Printf("b type %T b = % v \n",b3,b3)

	//func ParseInt(s string, base int, bitSize int) (i int64, err error)
	//返回字符串表示的整数值，接受正负号
	//base指定进制（2到36），如果base为0，则会从字符串前置判断，"0x"是16进制，"0"是8进制，否则是10进制；
	//bitSize指定结果必须能无溢出赋值的整数类型，0、8、16、32、64 分别代表 int、int8、int16、int32、int6
	var str2 string = "13445"
	b4 , _ := strconv.ParseInt(str2,10,64)
	fmt.Printf("b4 type %T b4 = % v \n",b4,b4)

	//func ParseFloat(s string, bitSize int) (f float64, err error)
	//解析一个表示浮点数的字符串并返回其值。
	
	var str3  string = "123455.1111"
	b5,  _  := strconv.ParseFloat(str3 , 32)
	fmt.Printf("b4 type %T b4 = % v \n",b5,b5)

	//放回值 都是int64 float64 想得到int32 float32
	var b6 = float32(b5)
	fmt.Printf("b6 type %T b6 = % v \n",b6,b6)

	//注意事项
	//在将String类型转成基本数据类型时，要确保String类型能够转成有效的数据，
	//比如我们可以把"123",转成一个整数，但是不能把"hello"转成一个整数，
	//如果这样做，Golang直接将其转成0，其它类型也是一样的道理.float=>0bool=>false

	b7, _ := strconv.ParseInt(str1,10,64)
	fmt.Printf("b7 type %T b7 = % v \n",b7,b7) // b7 type int64 b7 =  0

}