package main

import "fmt"

func main() {
	// 右辺の値により、型推論される
	// 右辺が変数なら、それと同じ型になる
	v := 42 // change me!
	fmt.Printf("v is of type %T\n", v)
}
