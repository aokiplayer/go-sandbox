package main

import "fmt"

type Vertex struct {
	Lat, Long float64
}

// map[キーの型]値の型 で宣言かな？
// この状態だと、 nil マップ（ゼロ値にあたり、キーを持ってないし追加できない）
var m map[string]Vertex

func main() {
	m = make(map[string]Vertex)
	m["Bell Labs"] = Vertex{
		40.68433, -74.39967,
	}
	fmt.Println(m["Bell Labs"])
}
