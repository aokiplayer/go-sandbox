package main

import "fmt"

func main() {
	// 初期値を設定しないスライスは nil となる。元となる配列がないから
	var s []int
	fmt.Println(s, len(s), cap(s))
	if s == nil {
		fmt.Println("nil!")
	}
}
