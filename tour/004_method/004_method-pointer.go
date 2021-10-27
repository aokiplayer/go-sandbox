package main

import (
	"fmt"
	"math"
)

type Vertex struct {
	X, Y float64
}

// このメソッドは変数レシーバを持ってるので、レシーバのコピーに対する操作
func (v Vertex) Abs() float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

// このメソッドはポインタレシーバを持ってるので、レシーバの値そのものを操作できる
// 他言語のクラスインスタンスに対するメソッドと同様の動きになる
func (v *Vertex) Scale(f float64) {
	v.X = v.X * f
	v.Y = v.Y * f
}

func main() {
	v := Vertex{3, 4}
	v.Scale(10)
	fmt.Println(v.Abs())
}
