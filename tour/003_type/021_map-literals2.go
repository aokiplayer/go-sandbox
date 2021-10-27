package main

import "fmt"

type Vertex struct {
	Lat, Long float64
}

var m = map[string]Vertex{
	// 本来なら "Bell Labs": Vertex{40.68433, -74.39967}, と書くところだけど、 map の値の型が単純なので省略できる
	// 簡単じゃないのは、むしろどんな？
	"Bell Labs": {40.68433, -74.39967},
	"Google":    {37.42202, -122.08408},
}

func main() {
	fmt.Println(m)
}
