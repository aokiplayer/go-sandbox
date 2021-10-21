package main

import "fmt"

type Vertex struct {
	X, Y int
}

var (
	v1 = Vertex{1, 2}

	// 一部のフィールドだけを指定することも可能
	v2 = Vertex{X: 1}
	v3 = Vertex{}

	// 新しく割り当てられた構造体へのポインタを取得。 *Vertex 型
	// しかし、この構文は何がうれしいんだろう？
	p  = &Vertex{1, 2}
)

func main() {
	fmt.Println(v1, p, v2, v3)
}
