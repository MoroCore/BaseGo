package main

import "fmt"

type Person struct {
	Num  int
	Name string
}

func (p Person) test() {
	fmt.Println(p.Num)
}
func (p Person) speak() {
	fmt.Println(p.Name, "是一个goodman~")
}
func (p Person) jisuan() {
	res := 0
	for i := 0; i <= 1000; i++ {
		res += i
	}
	fmt.Println(p.Name, "计算的结果是=", res)
}
func (p Person) jisuan2(n int) {
	res := 0
	for i := 0; i <= n; i++ {
		res += i
	}
	fmt.Println(p.Name, "计算的结果是=", res)
}
func (p Person) getSun(number1 int, number2 int) int {
	return number1 + number2
}

//结构体绑定方法
func main() {
	person := Person{}

	person.Name = "Tom"
	person.Num = 100

	person.test()
	person.speak()
	person.jisuan()
	person.jisuan2(10)
	person.getSun(10, 10)
}
