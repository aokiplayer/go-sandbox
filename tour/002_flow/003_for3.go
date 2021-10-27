package main

import "fmt"

func main() {
	sum := 1

	// 他の言語の while は、 for で記述できる
	// 初期化と後処理を区切るセミコロンは不要
	for sum < 1000 {
		sum += sum
	}
	fmt.Println(sum)
}
