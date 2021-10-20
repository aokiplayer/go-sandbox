package main

import "fmt"

// 引数の型は、同じのが続いたら最後にだけ書ける
// それ以外の引数が前や後にあった場合でも OK なのかな？
func add(x, y int) int {
	return x + y
}

func main() {
	fmt.Println(add(42, 13))
}
