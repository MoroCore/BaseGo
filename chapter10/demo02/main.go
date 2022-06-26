package main

import (
	"encoding/json"
	"fmt"
)

type Monster struct {
	Name  string `json:"name"`
	Age   int    `json:"age"`
	Skill string `json:"skill"`
}

//序列化 和 反序列化
func main() {

	//1：创建Monster变量
	monster := Monster{"牛魔王", 500, "芭蕉扇"}

	//2：将monster变量序列化为json格式字符串
	jsonStr, err := json.Marshal(monster)

	if err != nil {
		fmt.Println("json 处理 错误 ", err)
	}
	fmt.Println("jsonStr", string(jsonStr))
}
