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
type TV2 struct {
	*Goods
	*Brand
}

func main() {

	//嵌套匿名结构体后，也可以在创建结构体变量(实例)时，直接指定各个匿名结构体字段的值
	tv1 := TV{Goods{"电视台001", 5000.99}, Brand{"海尔", "山东"}}
	tv2 := TV{
		Goods{
			Price: 5000.99,
			Name:  "电视台001",
		},
		Brand{
			Name:    "海尔",
			Address: "山东",
		},
	}
	fmt.Println("tv1:", tv1)
	fmt.Println("tv2:", tv2)

	tv3 := TV2{&Goods{"电视台001", 5000.99}, &Brand{"海尔", "山东"}}
	tv4 := TV2{
		&Goods{
			Price: 5000.99,
			Name:  "电视台001",
		},
		&Brand{
			Name:    "海尔",
			Address: "山东",
		},
	}
	fmt.Println("tv3:", tv3)
	fmt.Println("tv4:", tv4)

}
