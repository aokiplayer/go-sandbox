package main

import (
	"fmt"
	"math/cmplx"
)

// 変数の宣言も、まとめてできる
var (
	ToBe   bool       = false
	// 符号なし int 型の 1 を 64 bit 左シフトして 1 を引く（uint で表せる最大値）
	MaxInt uint64     = 1<<64 - 1

	z      complex128 = cmplx.Sqrt(-5 + 12i)
)

func main() {
	fmt.Printf("Type: %T Value: %v\n", ToBe, ToBe)
	fmt.Printf("Type: %T Value: %v\n", MaxInt, MaxInt)
	fmt.Printf("Type: %T Value: %v\n", z, z)
}
