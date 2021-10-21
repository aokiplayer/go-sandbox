package main

import "fmt"

// 数値型で型のない定数は、状況によって必要な型を取る
// Big は 1 を 100 bit 左にずらしてるので、すごく大きい数値（これも定数としては表現できちゃう）
const (
	Big   = 1 << 100
	Small = Big >> 99
)

func needInt(x int) int {
	return x*10 + 1
}

func needFloat(x float64) float64 {
	return x * 0.1
}

func main() {
	//fmt.Println(Big)	// Println の引数は int だけど、 Big は int では表現しきれないのでエラー
	fmt.Println(Small)

	fmt.Println(needInt(Small))
	fmt.Println(needFloat(Small))
	fmt.Println(needFloat(Big))
}
