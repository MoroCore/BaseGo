package main

import "fmt"

type Studnet struct {
	name   string
	gender string
	age    int
	id     int
	score  float64
}

func (student *Studnet) say() string {
	infoStr := fmt.Sprintf("student的信息 name=[%v] gender=[%v], age=[%v] id=[%v] score=[%v]",
		student.name, student.gender, student.age, student.id, student.score)
	return infoStr
}

type Box struct {
	len    float64
	width  float64
	heigth float64
}

func (box *Box) getVolumn() float64 {
	return box.heigth * box.len * box.width
}
func main() {

	stu := Studnet{
		name:   "Tom",
		gender: "male",
		age:    18,
		id:     1000,
		score:  99.98,
	}
	fmt.Println(stu.say())
	box := Box{
		width:  10,
		len:    10,
		heigth: 10,
	}
	valumn := box.getVolumn()
	fmt.Println("valumn:", valumn)

}
