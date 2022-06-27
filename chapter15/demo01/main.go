package main

import (
	"encoding/json"
	"fmt"
)

/*
 * @GoName main
 * @Author Crow
 * @Email 648960069@qq.com
 * @Date 2022 9:09
 */

type Monster struct {
	Name string  `json:"monsterName"`
	Age  int     `json:"monsterAge"`
	Sal  float64 `json:"monsterSal"`
	Sex  string  `json:"monsterSex"`
}

func main() {

	m := Monster{
		Name: "玉兔精",
		Age:  20,
		Sal:  5555.33,
		Sex:  "female",
	}
	//反射的引出 为什么序列化的key是结构体的tag值 而不是字段的名称
	//反射  获取运行时 状态值
	marshal, _ := json.Marshal(m)
	fmt.Println("json result : ", string(marshal))
}
