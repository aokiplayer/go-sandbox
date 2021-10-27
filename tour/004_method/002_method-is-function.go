package main

import (
	"fmt"
	"math"
)

type Vertex struct {
	X, Y float64
}

// 前の例をメソッドではなく関数として書くとこうなる
func Abs(v Vertex) float64 {
	return math.Sqrt(v.x*v.x + v.Y*v.Y)
}

func main() {
	v := Vertex{3, 4}
	fmt.Println(Abs(v))
}
