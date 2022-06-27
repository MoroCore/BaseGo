package main

import (
	"fmt"
	"github.com/garyburd/redigo/redis"
)

/*
 * @GoName main
 * @Author Crow
 * @Email 648960069@qq.com
 * @Date 2022 15:08
 */

func main() {

	redisCon, err := redis.Dial("tcp", "192.168.10.10:6379")
	if err != nil {
		fmt.Println("redis.Dirl err = ", err)
		return
	}
	defer redisCon.Close()
	_, err = redisCon.Do("Set", "name", "moro")
	if err != nil {
		fmt.Println("Set err = ", err)
		return
	}
	name, err := redis.String(redisCon.Do("Get", "name"))
	if err != nil {
		fmt.Println("set err = ", err)
		return
	}
	fmt.Println("操作 ok", name)
}
