package main

type A struct {
	name string
}
type B struct {
	a A //有名结构体
}

func main() {
	//如果中是一个有名结构体，则访问有名结构体的字段时，就必须带上有名结构体的名字
	//比如d.a.Name
	var b B
	b.a.name = "jack"
}
