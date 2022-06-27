package main

import (
	"fmt"
	"reflect"
)

/*
 * @GoName main
 * @Author Crow
 * @Email 648960069@qq.com
 * @Date 2022 10:52
 */

func main() {
	var a Monster = Monster{
		Name:  "moro",
		Age:   29,
		Score: 30.3,
		Sex:   "carlos",
	}
	testStruct(a)
}

type Monster struct {
	Name  string  `json:"name"`
	Age   int     `json:"age"`
	Score float32 `json:"score"`
	Sex   string
}

func (s Monster) Set(name string, age int, score float32, sex string) {
	s.Name = name
	s.Age = age
	s.Score = score
	s.Sex = sex
}
func (s Monster) Show() {
	fmt.Println("---start----")
	fmt.Println(s)
	fmt.Println("---end-----")
}
func testStruct(a interface{}) {

	typ := reflect.TypeOf(a)
	val := reflect.ValueOf(a)

	kind := val.Kind()

	if kind != reflect.Struct {
		return
	}
	numFiled := val.NumField()

	for i := 0; i < numFiled; i++ {
		fmt.Printf("Filed %d : 值为 = %v \n", i+1, val.Field(i))
		tag := typ.Field(i).Tag.Get("json")
		if tag != "" {
			fmt.Printf("Field %d: tag 为=%v\n", i, tag)
		}
	}
	numMethod := val.NumMethod()
	fmt.Printf("struct has %d methods\n", numMethod)
	method := val.Method(1)
	fmt.Println(method.String())
	var params []reflect.Value
	val.Method(1).Call(params)
	//fmt.Println("res=", res[0].Int())
}
