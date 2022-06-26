package main

import "fmt"
// import "math"

//switch分支控制
func main(){
	//1)switch语句用于基于不同条件执行不同动作，每一个case分支都是唯一的，从上到下逐一测试，直到匹配为止。
	//2)匹配项后面也不需要再加break

	//1)switch的执行的流程是，先执行表达式，得到值，然后和case的表达式进行比较，如果相等，就匹配到，然后执行对应的case的语句块，然后退出switch控制。
	//2)如果switch的表达式的值没有和任何的case的表达式匹配成功，则执行default的语句块。执行后退出switch的控制.
	//3)golang的case后的表达式可以有多个，使用逗号间隔.
	//4)golang中的case语句块不需要写break,因为默认会有,即在默认情况下，当程序执行完case语句块后，就直接退出该switch控制结构。

	
}
//4switch快速入门案例
//请编写一个程序，该程序可以接收一个字符，比如:a,b,c,d,e,f,g据用户的输入显示相依的信息.要求使用switch语句完成
func test1(){
	
	var key byte
	fmt.Printf("请输入一个字符 a b c d e f g")
	fmt.Scanf("%c",&key)

	switch key {
		case 'a':
			fmt.Println("周一，猴子穿新衣")
		case 'b':
			fmt.Println("周二，猴子当小二")
		case 'c':
			fmt.Println("周三，猴子爬雪山")
		default:
			fmt.Println("输入有误")	
	}
	//细节
	// 1)case/switch后是一个表达式(即：常量值、变量、一个有返回值的函数等都可以
		// switch test(1)+1  
	// 2)case后的各个表达式的值的数据类型，必须和switch的表达式数据类型一致
		// n1 int32 = 20 
		// n2 int64 = 20
		//      switch n1
		//      	case n2  //类型不匹配
	// 3)case后面可以带多个表达式，使用逗号间隔。比如case表达式1,表达式2...
		//  case n2,10,4
	// 4)case后面的表达式如果是常量值(字面量)，则要求不能重复
	// 5)case后面不需要带break,程序匹配到一个case后就会执行对应的代码块，然后退出switch，如果一个都匹配不到，则执行default
	// 6)default语句不是必须的.
	// 7)switch后也可以不带表达式，类似if--else分支来使用。
	// 8)switch后也可以直接声明/定义一个变量，分号结束，不推荐。
	// 9)switch穿透-fallthrough，如果在case语句块后增加fallthrough,则会继续执行下一个case，也叫switch穿透
	// 10)TypeSwitch：switch语句还可以被用于type-switch来判断某个interface变量中实际指向的变量类型

	//总结了什么情况下使用switch,什么情况下使用if
	//1)如果判断的具体数值不多，而且符合整数、浮点数、字符、字符串这几种类型。建议使用swtich语句，简洁高效。
	//2)其他情况：对区间判断和结果为bool类型的判断，使用if，if的使用范围更广。
}