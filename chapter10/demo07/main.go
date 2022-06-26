package main

import (
	"GoLang/chapter10/model"
	"fmt"
)

func main() {
	//如果挎包定义的结构体 首写字母大写 可以直接使用
	stu := model.Studnet{
		// name:  "tom",
		// score: 78.9,
	}
	stu.SetScore(10)
	fmt.Println(stu.GetScore())
	fmt.Println(stu)
	//如果挎包定义的结构体首字母小写  通过工厂模式来解决
	user := model.NewUser("10", 11)
	fmt.Println("user:", &user)
}
