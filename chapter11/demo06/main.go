package main

import "fmt"

type Goods struct {
	Name  string
	Price float64
}
type Brand struct {
	Name    string
	Address string
}
type TV struct {
	Goods
	Brand
}

func main() {
	var tv TV
	//1）如嵌入的匿名结构体有相同的字段名或者方法名，则在访问时，需要通过匿名结构体类型名来区分
	fmt.Println(tv.Goods.Name)
	fmt.Println(tv.Price)
	// 2)为了保证代码的简洁性，建议大家尽量不使用多重继承
}
