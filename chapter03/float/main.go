package main

import "fmt"
func main(){

	//小数类型就是用于存放小数的，比如1.20.23-1.911

	//单精度 float32  float64
	
	//关于浮点数在机器中存放形式的简单说明,浮点数=符号位+指数位+尾数位说明：浮点数都是有符号的.

	var price float32 = 89.12
	fmt.Println("price=",price)
	var num1 float32 = -0.0000089
	var num2 float64 = -0.0000089
	fmt.Println("num1=",num1,"num2=",num2)

	//2)尾数部分可能丢失，造成精度损失。-123.0000901
	var num3 float32 =  -123.0000901
	var num4 float64 =  -123.0000901
	fmt.Println("num3=",num3,"num4=",num4)
	//float64的精度比float32的要准确.
	//如果我们要保存一个精度高的数，则应该选用float64

	//细节
	//1)Golang浮点类型有固定的范围和字段长度，不受具体OS(操作系统)的影响。
	//2)Golang的浮点型默认声明为float64类型。
	var num5 = 1.1
	fmt.Printf("num5的数据类型 %T \n",num5)

	// 3)浮点型常量有两种表示形式
	//十进制数形式：如：5.12.512(必须有小数点）
	//科学计数法形式:如：5.1234e2=5.12*10的2次方   5.12E-2=5.12/10的2次方
	//十进制数形式
	num6 := 5.12
	num7 := 0.123
	fmt.Println("num6=",num6,"num7=",num7)

	num8 := 5.22e2
	num9 := 66e-2
	fmt.Println("num8=",num8,"num9=",num9)

}