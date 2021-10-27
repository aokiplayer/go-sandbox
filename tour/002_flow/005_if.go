package main

import (
	"fmt"
	"math"
)

func sqrt(x float64) string {
	if x < 0 {
		return sqrt(-x) + "i"
	}
	// Sprint ってなんだろう？
	// 引数の値をフォーマットした文字列を返す。 Println を標準出力じゃなくて文字列で返す感じ？
	return fmt.Sprint(math.Sqrt(x))
}

func main() {
	fmt.Println(sqrt(2), sqrt(-4))
}
