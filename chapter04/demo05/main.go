package main

import (
	"fmt"
)

//键盘输入语句
func main(){

	//在编程中，需要接收用户输入的数据，就可以使用键盘输入语句来获取


	//步骤
	//	1)导入fmt包
	//	2)调用fmt包的fmt.Scanln()或者fmt.Scanf()

	// func Scanln(a ...interface{}) (n int, err error)  a ...interface{}  
	// Scanln类似Scan，但会在换行时才停止扫描。最后一个条目后必须有换行或者到达结束位置。
	var name string
	var age byte
	var sal float32
	var isPass bool

	// fmt.Println("请输入名字 ")
	// fmt.Scanln(&name)
	// fmt.Println("请输入年龄 ")
	// fmt.Scanln(&age)
	// fmt.Println("请输入薪水 ")
	// fmt.Scanln(&sal)
	// fmt.Println("请输入是否通过考试 ")
	// fmt.Scanln(&isPass)

	fmt.Printf("名字是 %v \n 年龄是 %v \n 薪水是%v \n 是否通过考试 %v \n",name,age,sal,isPass)


	//func Sscanf(str string, format string, a ...interface{}) (n int, err error)
	//Sscanf从字符串str扫描文本，根据format 参数指定的格式将成功读取的空白分隔的值保存进成功传递给本函数的参数。返回成功扫描的条目个数和遇到的任何错误。
	fmt.Println("”请输入你的姓名，年龄，薪水，是否通过考试， 使用空格隔开")
	fmt.Scanf("%s %d %f %t",&name,&age,&sal,&isPass)
	fmt.Printf("名字是 %v \n 年龄是 %v \n 薪水是%v \n 是否通过考试 %v \n",name,age,sal,isPass)

}