package main

import "fmt"

// bool 型の初期値は false？
// 変数のスコープはパッケージ内と関数内の 2 つだけ？
var c, python, java bool

func main() {
	// int 方の初期値は 0？
	var i int
	fmt.Println(i, c, python, java)
}
