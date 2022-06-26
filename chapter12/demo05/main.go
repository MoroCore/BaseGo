package main

import (
	"encoding/json"
	"fmt"
)

/**
 * @GoName demo05
 * @Author Crow
 * @Email 648960069@qq.com
 * @Date 2022/6/25 16:12
 */
func main() {
	testStruce()
	testMap()
	testSlice()
	testFloat()
}

//将结构体 map 切片序列化
type Monster struct {
	Name     string
	Age      int
	Birthday string
	Sal      float64
	Skill    string
}

func testStruce() {

	monster := Monster{
		Name:     "牛魔王",
		Age:      500,
		Birthday: "2011-11-11",
		Sal:      80000.0,
		Skill:    "牛魔拳",
	}
	marshal, err := json.Marshal(monster)
	if err != nil {
		fmt.Printf("序列化失败 err = %v", err)
	}
	fmt.Printf("monster序列化后 = %v \n", string(marshal))
}
func testMap() {
	var a map[string]interface{}
	a = make(map[string]interface{})
	a["name"] = "红孩儿"
	a["age"] = 30
	a["adress"] = "beijing"
	marshal, err := json.Marshal(a)
	if err != nil {
		fmt.Printf("序列化失败 err = %v", err)
	}
	fmt.Printf("map序列化后 = %v \n", string(marshal))
}

// 对切片进行
func testSlice() {

	var slice []map[string]interface{}

	var m1 = make(map[string]interface{})
	m1["name"] = "jack"
	m1["age"] = 17
	m1["adress"] = "北京"
	slice = append(slice, m1)
	var m2 = make(map[string]interface{})
	m2["name"] = "moro"
	m2["age"] = 17
	m2["adress"] = "上海"
	slice = append(slice, m2)
	marshal, err := json.Marshal(slice)
	if err != nil {
		fmt.Printf("序列化失败 err = %v", err)
	}
	fmt.Printf("slice序列化后 = %v \n", string(marshal))
}
func testFloat() {

	num := 1234.11
	marshal, err := json.Marshal(num)
	if err != nil {
		fmt.Printf("序列化失败 err = %v", err)
	}
	fmt.Printf("slice序列化后 = %v \n", string(marshal))
}
