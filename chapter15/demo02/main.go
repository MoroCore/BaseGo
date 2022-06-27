package main

/*
 * @GoName main
 * @Author Crow
 * @Email 648960069@qq.com
 * @Date 2022 9:25
 */

func main() {

	test1 := func(v1 int, v2 int) {
	}
	test2 := func(v1 int, v2 int) {

	}
	bridge := func(call interface{}, args ...interface{}) {

	}
	bridge(test1, 1, 2)
	bridge(test2, 2, 4)
}
