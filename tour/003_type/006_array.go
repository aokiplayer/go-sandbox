package main

import "fmt"

func main() {
	// 配列は型の書き方が [n]要素の型 であること以外は、他の言語とそんなに変わらない
	var a [2]string
	a[0] = "Hello"
	a[1] = "World"
	fmt.Println(a[0], a[1])
	fmt.Println(a)

	primes := [6]int{2, 3, 5, 7, 11, 13}
	fmt.Println(primes)
}
