package main

import "fmt"

type AInterface interface {
	Say()
}
type BInterface interface {
	Hello()
}
type Monster struct {
}

func (m Monster) Hello() {
	fmt.Println("Monster Hello()~~~~~")
}
func (m Monster) Say() {
	fmt.Println("Monster Say()~~~~~")
}
func main() {
	//Monster实现了AInterface和 BInterface
	var monster Monster
	var a AInterface = monster
	var b BInterface = monster
	a.Say()
	b.Hello()
}
