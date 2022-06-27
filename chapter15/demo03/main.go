package main

import (
	"fmt"
	"reflect"
)

/*
 * @GoName main
 * @Author Crow
 * @Email 648960069@qq.com
 * @Date 2022 10:21
 */

func main() {

	//var num int = 100

	//reflectTest01(num)

	stu := Student{
		Name: "tom",
		Age:  12,
	}
	reflectTest2(stu)
}

//测试基本数据类型
func reflectTest01(b interface{}) {

	tType := reflect.TypeOf(b)
	fmt.Println("tType=", tType)

	valueOf := reflect.ValueOf(b)
	fmt.Println("valueOf=", valueOf)

	n2 := valueOf.Int() + 2
	fmt.Println("n2=", n2)

	fmt.Printf("valueOf = %v valueOf type = %T \n ", valueOf, valueOf)

	iV := valueOf.Interface()
	num2 := iV.(int)
	fmt.Println("num2=", num2)
	fmt.Printf("num2 type = %T \n", num2)
}

//测试结构体的反射
func reflectTest2(b interface{}) {

	//type =  *reflect.rtype
	//typeOf := reflect.TypeOf(b)
	//value = reflect.Value
	valueOf := reflect.ValueOf(b)
	//fmt.Printf("typeOf type = %T \n valueOf type = %T", typeOf, valueOf)

	value := valueOf.Interface()
	student := value.(Student)
	fmt.Printf("student type = %T  \n student = %v \n", student, student)
}

type Student struct {
	Name string
	Age  int
}
