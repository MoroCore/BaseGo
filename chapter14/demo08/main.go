package main

import "fmt"

/*
 * @GoName main
 * @Author Crow
 * @Email 648960069@qq.com
 * @Date 2022 11:55
 */

type Cat struct {
	Name string
	Age  int
}

func main() {

	//存放任意类似的时候出现的问题

	allChan := make(chan interface{}, 10) //该管道可以存放任意类型

	allChan <- Cat{
		Name: "Moro",
		Age:  12,
	} //Cat类型在存放到 allChan中时, Cat类型被转换为了interface{}类型

	cat := <-allChan
	//fmt.Println(cat.Name) error
	c := cat.(Cat) //加个类型断言 将interface 转换为Cat
	fmt.Println(c.Age)
}
