package main

import "fmt"

//结构体内存的分配机制
type Person struct {
	Name string
	Age  int
}

func main() {

	//结构体在内存中是单独存在的
	var person1 Person

	person1.Name = "Moro"
	person1.Age = 18

	var person2 *Person = &person1

	// person2.Name = "Carlos"
	// person2.Age = 20

	fmt.Printf("persion1的地址:%p \n", &person1)
	fmt.Printf("persion2的地址:%p", person2)
}
