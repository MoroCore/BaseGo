package main

import (
	"fmt"
)

//算术运算符
func main(){

	//算术运算符是对数值类型的变量进行运算的，比如：加减乘除。在Go程序中使用的非常多
	//+,-,*,/,%,++,--,重点讲解/、%

	//  /  说明，如果运算的数都是容数那么除后，去控小数部分保留教数部分
	fmt.Println(10/-3) //2
	// 如果我们希望保留小数部分，则需要有浮点数参与运算
	fmt.Println(10.0/4) //2.5

	//%的使用特点
	fmt.Println("10%3=",10%3)//=1
	fmt.Println("-10%3=",-10%3)//=-10-(-10)/3*3=-10-(-9)=-1
	fmt.Println("10%-3=",10%-3)//=1  10-10/(-3)*(-3)
	fmt.Println("-10%-3=",-10%-3)//=-1

	//++和--的使用
	var i int=10
	i++ //等价1=1+1
	fmt.Println("i=", i) // 11  unexpected ++, expecting  如果在打印函数中使用 i++
	i-- //等价1=1-1
	fmt.Println("i=", i) // 10

	//3算术运算符使用的注意事项
	//	1)对于除号"/"，它的整数除和小数除是有区别的：整数之间做除法时，只保留整数部分而舍弃小数部分。例如：x:=19/5,结果是3
	//	2)当对一个数取模时，可以等价a%b=a-a/b*b，这样我们可以看到取模的一个本质运算。
	//	3)Golang的自增自减只能当做一个独立语言使用时，不能这样使用
	//	4)Golang的++和--只能写在变量的后面，不能写在变量的前面，即：只有a++a--没有++a
	//	Golang的设计者去掉c/java中的自增自减的容易混淆的写法，让Golang更加简洁，统一。(强制性的)

	week,day := getWeekAndWhatDayByDays(8)
	fmt. Printf( "%d个星期零%d天\n", week, day)

	KWenDu :=getNewKWengDu(134.2)
	fmt. Printf( "华氏温度%v,对应的摄氏度%v", 134.2, KWenDu)
}

//1)假如还有97天放假，问：xx个星期零xx天
func getWeekAndWhatDayByDays(days int) (int,int){

	week := days / 7
	day := days % 7

	return week , day
}
//2)定义一个变量保存华氏温度，华氏温度转换摄氏温度的公式为：5/9*(华氏温度-100),请求出华氏温度对应的摄氏温度
func getNewKWengDu(Huashi float32) float32{
	
	return 5.0/9*(Huashi-100)
}