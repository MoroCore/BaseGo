package main

import "fmt"
import "math"

//分支控制
func main(){

	//1)单分支
	//	编写一个程序,可以输入人的年龄,如果该同志的年龄大于18岁,则输出"你年龄大于18,要对自己的行为负责!
	// var age int

	// fmt.Println("请输入你的年龄")
	// fmt.Scanln(&age)

	// if age > 18 {
	// 	fmt.Println("你年龄大于18,要对自己的行为负责!")
	// }
	////golang支持在if中，直接定义-一个变量，比如下面
	// if age := 20 ; age > 18 {
	// 	fmt.Println("你年龄大于18,要对自己的行为负责!")
	// }

	//双分支控制
	//编写一个程序,可以输入人的年龄,如果该同志的年龄大于18岁,则输出“你年龄大于18,要对自己的行为负责!”。否则,输出”你的年龄不大这次放过你了.”

	var age int
	fmt.Println("请输入你的年龄")
	fmt.Scanln(&age)

	if age > 18 {
		fmt.Println("你年龄大于18,要对自己的行为负责!")
	} else {
		fmt.Println("你的年龄不大这次放过你了")
	}
	//细节  {} 必须添加  else 不能换行
	test1()

	//多分支控制
	/* 
		1)多分支的判断流程如下:
			(1)先判断条件表达式1是否成立，如果为真，就执行代码块1
			(2)如果条件表达式1如果为假，就去判断条件表达式2是否成立，如果条件表达式2为真，就执行代码块2
			(3)依次类推.
			(4)如果所有的条件表达式不成立，则执行else的语句块。
		2)else不是必须的。
		3)多分支只能有一个执行入口。
	*/

}

//编写程序，声明2个int32型变量并赋值。判断两数之和，如果大于等于50，打印“helloworld!
func test1(){
	var n1 = 10
	var n2 = 20
	if n1 + n2 <= 50 {
		fmt.Println("Hello,World!")
	}
}
//6)编写程序，声明2个float64型变量并赋值。判断第一个数大于10.0，且第2个数小于20.0，打印两数之和。
func test2(){
	
	var n3 float64 = 11.0
	var n4 float64 = 12.0

	if n3 > 10.0 && n4 < 20.0 {
		fmt.Println("和=",(n3+n4))
	}
}

//定义两个变量int32，判断二者的和，是否能被3又能被5整除，打印提示信息
func test4(){

	var n1 = 10
	var n2 = 20

	if (n1 + n2 ) % 3 == 0 &&  (n1 + n2 ) % 5 == 0 {
		fmt.Println("'能被3又能被5整除")
	} 
}
//判断一个年份是否是闰年，闰年的条件是符合下面二者之一：(1)年份能被4整除，但不能被100整除；(2)能被400整除
func test5(){

	var year int = 2019
	if (year % 4 == 0  && year % 100 != 0) || year % 400 == 0 {
		fmt. Println(year,"是润年~")
	} 
}

//岳小鹏参加Golang考试，他和父亲岳不群达成承诺：
//如果：成绩为100分时，奖励一辆BMW；
//成绩为(80，99]时，奖励一台iphone7plus；
//当成绩为[60,80]时，奖励一个iPad；
//其它时，什么奖励也没有。
//请从键盘输入岳小鹏的期末成绩，并加以判断
func test6(){

	var source int;

	//细节  当判断为true时 程序不会在向下判断
	fmt.Scanln(&source)

	if source == 100 {
		fmt.Println("奖励一辆BMW")
	}else if source > 80 && source <= 99{
		fmt.Println("奖励一台iphone7plus")
	}else if source >=60 && source <= 80 {
		fmt.Println("奖励一个iPad")
	}else {
		fmt.Println("什么奖励也没有")
	}
}

//求ax2 +bx+c=0方程的根。a,b,c分别为函数的参数，如果: b2-4ac>0， 则有两个解;
//b2 -4ac=0， 则有一个解; b2-4ac<0， 则无解;
func test7(){

	var a float64
	var b float64
	var c float64

	fmt.Scanf("%d %d %d",&a,&b,&c)

	m := b * b - 4 * a * c
	
	if m > 0 {
		x1 := (-b + math.Sqrt(m)) / 2 * a
		x2 := (-b - math.Sqrt(m)) / 2 * a
		fmt.Printf("x1 = %v  x2 = %v ",x1,x2)
	}else if m == 0 {
		x1 := (-b + math.Sqrt(m)) / 2 * a
		fmt.Printf("x1 = %v",x1)
	}else {
		fmt.Printf("无解")
	}

}
//大家都知道，男大当婚，女大当嫁。那么女方家长要嫁女儿，当然要提出一定的条件:高:
//180cm以上;富:财富1千万以上;帅:是。条件从控制台输入。
//1)如果这三个条件同时满足，则:“我一-定要嫁给他!!!”
//2) 如果三个条件有为真的情况，则:“嫁吧，比上不足，比下有余。”
//3) 如果三个条件都不满足，则:“不嫁!
func test8() {
	var money float64
	var height int
	var handsome bool

	fmt.Println("请输入身高(厘米)")
	fmt.Scanln(&height)
	fmt.Println("请输财富(千万)")
	fmt.Scanln(&money)
	fmt.Println("请输入是否帅(true/false))")
	fmt.Scanln(&handsome)

	if height > 180 && money > 1.0 && handsome{
		fmt.Println("我一定要嫁给他！！！")
	}else if height > 180 || money > 1.0 || handsome {
		fmt.Printf("嫁吧，比上不足，比下有余。")
	}else {
		fmt.Printf("不嫁!......")
	}
}