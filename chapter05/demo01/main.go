package main

import "fmt"

//顺序控制
func main(){

	//程序从上到下逐行地执行，中间没有任何判断和跳转。
	//一个案例说明，必须下面的代码中，没有判断，也没有跳转.因此程序按照默认的流程执行，即顺序控制。
	getWeekAndWhatDayByDays(97)
}
//1)假如还有97天放假，问：xx个星期零xx天
func getWeekAndWhatDayByDays(days int) (int,int){

	week := days / 7
	day := days % 7

	


	// Golang中定义变量时采用合法的前向引用。如
	var num1 int = 10//声明了num1
	var num2 int = num1 + 20//使用num1fmt.Println(num2)
	fmt.Println(num2)
	return week , day
	// 错误形式：
	// var num3 int = num4 + 20//使用num4
	// var num4 int = 10//声明num1(×)
	// fmt.Println(num2)
}