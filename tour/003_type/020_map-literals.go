package main

import "fmt"

type Vertex struct {
	Lat, Long float64
}

// make 関数を使わずに、リテラルを指定して初期化できるのかな？
var m = map[string]Vertex{
	"Bell Labs": Vertex{
		40.68433, -74.39967,
	},
	"Google": Vertex{
		37.42202, -122.08408,
	},
}

func main() {
	fmt.Println(m)
}
