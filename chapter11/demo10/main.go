package main

type Hero struct {
	Name string
	Age  int
}

type HeroSlice []Hero

func (hs HeroSlice) Len() int {
	return len(hs)
}

func (hs HeroSlice) Less(i int, j int) bool {
	return hs[i].Age < hs[j].Age
}

func (hs HeroSlice) Swap(i int, j int) {
	hs[i], hs[j] = hs[j], hs[i]
}

type Studnet struct {
	Name  string
	Age   int
	Score float64
}
