package main

import (
	"fmt"
	"math"
)

// 別パッケージの型にメソッドを定義できないので、型を作ってる
// この構文（Defined type）は Type alias と異なるので注意
//type MyFloat = float64	// こっちは Type alias
type MyFloat float64

func (f MyFloat) Abs() float64 {
	if f < 0 {
		// MyFloat と float64 は別の型なので、型変換が必要
		return float64(-f)
	}
	return float64(f)
}

func main() {
	f := MyFloat(-math.Sqrt2)
	fmt.Println(f.Abs())
}
