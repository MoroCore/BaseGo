package model

//定义一个结构体
type Studnet struct {
	name  string
	score float64
}

type user struct {
	Name  string
	Score float64
}

func NewUser(n string, s float64) *user {
	return &user{
		Name:  n,
		Score: s,
	}
}
func (s *Studnet) GetScore() float64 {
	return s.score
}
func (s *Studnet) SetScore(score float64) {
	s.score = score
}
