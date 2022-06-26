package main

func main() {

	//声明map切片 方式1
	var monsters []map[string]string
	monsters = make([]map[string]string, 2)
	//
	if monsters[0] == nil {
		monsters[0] = make(map[string]string)
		monsters[0]["name"] = "moro"
		monsters[0]["age"] = "3000"
	}
	if monsters[1] == nil {
		monsters[1] = make(map[string]string)
		monsters[1]["name"] = "moro"
		monsters[1]["age"] = "3000"
	}
	// if monsters[2] == nil {
	// 	monsters[2] = make(map[string]string)
	// 	monsters[2]["name"] = "moro"
	// 	monsters[2]["age"] = "3000"
	// }
	newMonster := map[string]string{
		"name": "1",
	}
	monsters = append(monsters, newMonster)
}
