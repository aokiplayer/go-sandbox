package main

import "fmt"

// Go では型宣言は後ろに書くスタイル
func add(x int, y int) int {
	return x + y
}

func main() {
	fmt.Println(add(42, 13))
}
