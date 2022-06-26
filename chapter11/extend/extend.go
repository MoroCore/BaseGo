package extend

type Goods struct {
	Name  string
	Price int
}

type Book struct {
	Goods //这里就是嵌套匿名结构体Goods
	Write string
}
