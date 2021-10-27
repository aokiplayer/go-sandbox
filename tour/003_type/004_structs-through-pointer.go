package main

import "fmt"

type Vertex struct {
	X int
	Y int
}

func main() {
	v := Vertex{1, 2}
	p := &v // v へのポインタ p
	p.X = 1e9	// 本来は (*p).X と書くところだけど、簡易的にこう書ける
	fmt.Println(v)
}
