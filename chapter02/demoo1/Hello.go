package main	//Go 语言的一个文件都要归属于一个包
				//main函数 必须在 main包
import "fmt"	//导入包

func main(){	//main 是函数名，是一个主函数，即我们程序的入口。

	fmt.Println("Hello World")

	// 5 Go 编译器是一行行进行编译的，因此我们一行就写一条语句，不能把多条语句写在同一个，否则报错
	// fmt.Println("Hello1") fmt.Println("Hello2")

	// 6) go 语言定义的变量或者 import 的包如果没有使用到，代码不能编译通过。
	//num declared but not usedg
	// var num = 10

	// 7 转移字符  
	// \t : 表示一个制表符，通常使用它可以排版
	fmt.Println("Carlos \t Moro")
	// \n ：换行符
	fmt.Println("Carlos \n Moro")
	//  \\ ：一个\
	fmt.Println("Carlos \\ Moro")
	//  \" ：一个"
	fmt.Println("Carlos \" Moro")
	//\r ：一个回车 fmt.Println("天龙八部雪山飞狐\r 张飞");
	// 从当 前行的最前面开始输出，覆盖掉以前内容
	fmt.Println("天龙八部雪山飞狐\r 张飞")

	//12 注释
	// 行注释   块注释
	// 语法   //     /*  */
	// go官方推荐使用 行注释    
	// 细节 对于行注释和块注释，被注释的文字，不会被Go编译器执行 
	//块注释里面不允许有块注释嵌套

	/*         块注释
	   ...........................................
	*/

	// Go设计者思想:一个问题尽量只有一个解决方法

	// 一行最长不超过80个字符，超过的请使用换行展示，尽量保持格式优雅
	//  + 在结尾
	fmt.Println("aaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaa"+
			"bbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbb"+
			"ccccccccccccccccccccccccccccc")
}