package main

import "fmt"

func main() {
	// make 関数でスライスを作成できる
	a := make([]int, 5)
	printSlice("a", a)	// length も capacity も 5

	b := make([]int, 0, 5)	// length: 0, capacity: 5
	printSlice("b", b)

	c := b[:2]
	printSlice("c", c)

	d := c[2:5]
	printSlice("d", d)
}

func printSlice(s string, x []int) {
	fmt.Printf("%s len=%d cap=%d %v\n",
		s, len(x), cap(x), x)
}
