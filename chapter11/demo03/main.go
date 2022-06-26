package main

type A struct {
	Name string
	age  int
}
type B struct {
	Name  string
	Score float64
}

type C struct {
	A
	B
}

func main() {

	var c C
	//如果c没有Name字段，而A和 B有Name,这时就必须通过指定匿名结构体名字来区分
	//否则c.Name就会包编译错误,这个规则对方法也是一样的!
	c.A.Name = "Tom"
	// c.Name = "T"

}
