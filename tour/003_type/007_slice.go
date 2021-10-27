package main

import "fmt"

func main() {
	primes := [6]int{2, 3, 5, 7, 11, 13}

	// 配列よりも slice の方が一般的。可変長の配列
	// 配列[インデックス1:インデックス2] の記法は、配列からスライスを作成（Swift の半開 Range に近い）
	var s []int = primes[1:4]
	fmt.Println(s)
}
