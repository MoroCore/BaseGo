package main

import (
	"fmt"
)

func main(){


	//在go中，数据类型都有一个默认值，当程序员没有赋值时，就会保留默认值，在go中，默认值又叫零值。
	
	//基本数据类型的默认值如下

	var a int //0
	var b float32 //0
	var c bool //false
	var d string //""
	fmt.Printf("a=%v,b=%v,c=%v,d=%v",a,b,c,d)
}