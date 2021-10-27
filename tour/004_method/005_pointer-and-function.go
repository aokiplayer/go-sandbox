package main

import (
	"fmt"
	"math"
)

type Vertex struct {
	X, Y float64
}

func Abs(v Vertex) float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

// こっちの関数は、構造体のアドレスを受け取ってるので操作できる
func Scale(v *Vertex, f float64) {
	v.X = v.X * f
	v.Y = v.Y * f
}

func main() {
	v := Vertex{3, 4}
	// ここでは、構造体のアドレスを渡している
	Scale(&v, 10)
	fmt.Println(Abs(v))
}
