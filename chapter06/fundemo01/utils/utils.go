package utils

import (
	"fmt"
)

//细节  将需要其他包调用的函数  函数名首字母大写 类型Java Public
func Cal(n1 float64, n2 float64, opterator byte) float64 {

	var res float64

	switch opterator {
		case '+':
			res = n1 + n2
		case '-':
			res = n1 - n2
		case '*':
			res = n1 * n2
		case '/':
			res = n1 / n2
		default:
			fmt.Println("操作符输入错误")
	}
	return res
}