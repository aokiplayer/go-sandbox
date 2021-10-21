package main

import (
	"fmt"
	"math"
)

func main() {
	x := 3
	y := 4
	// 明示的な型変換が必要（整数から小数であっても必要）
	f := math.Sqrt(float64(x*x + y*y))
	z := uint(f)
	fmt.Println(x, y, z)
}
