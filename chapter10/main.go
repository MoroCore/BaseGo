package main

import "fmt"

//定义一个结构体
type Cat struct {
	Name  string
	Age   int
	Color string
	Hobby string
}

//如果结构体的字段类型是:指针，slice，和map的零值都是nil，即还没有分配空间
//如果需要使用这样的字段,需要先make，才能使用.
type Person struct {
	Name   string
	Age    int
	Scores [5]float64
	prt    *int
	slice  []int
	map1   map[string]string
}

func main() {

	//创建一个Cat变量
	// var cat1 Cat
	// cat1.Name = "小白"
	// cat1.Age = 3
	// cat1.Color = "白色"
	// cat1.Hobby = "吃鱼"

	// fmt.Println(cat1)

	var person Person
	fmt.Println(person)

	if person.prt == nil {
		fmt.Println("ok1")
	}
	if person.slice == nil {
		fmt.Println("ok2")
	}
	if person.map1 == nil {
		fmt.Println("ok3")
	}
}
