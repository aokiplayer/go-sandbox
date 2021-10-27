package main

import "fmt"

func main() {
	// スライス自体は値を持ってない。ので、この構文は「配列を作ってそれを参照するスライス」の宣言ってことかな？
	q := []int{2, 3, 5, 7, 11, 13}
	fmt.Println(q)

	r := []bool{true, false, true, true, false, true}
	fmt.Println(r)

	// 前の部分で構造体の形を定義して、それのスライスを作成。後ろの部分では、その構造体の形の初期値を設定
	// 配列やスライスの宣言部分では [] なのに、初期値設定は {} なのがややこしい
	s := []struct {
		i int
		b bool
	}{
		{2, true},
		{3, false},
		{5, true},
		{7, true},
		{11, false},
		{13, true},
	}
	fmt.Println(s)
}
