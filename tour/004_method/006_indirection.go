package main

import "fmt"

type Vertex struct {
	X, Y float64
}

func (v *Vertex) Scale(f float64) {
	v.X = v.X * f
	v.Y = v.Y * f
}

// 構造体のポインタを引数として受け取る関数
func ScaleFunc(v *Vertex, f float64) {
	v.X = v.X * f
	v.Y = v.Y * f
}

func main() {
	v := Vertex{3, 4}
	// ここはレシーバがポインタじゃないように見えるけど、実際はポインタが渡される
	// 内部的にこうなる (&v).Scale(2)
	v.Scale(2)
	// 構造体のポインタを渡す
	ScaleFunc(&v, 10)

	p := &Vertex{4, 3}
	p.Scale(3)
	// 2 行上で構造体へのポインタを取得しているので渡してる
	ScaleFunc(p, 8)

	fmt.Println(v, p)
}
