package main

import "fmt"

func main() {
	// スライスには length と capacity がある
	// length は、今の要素数。 capacity は元の配列の要素数
	s := []int{2, 3, 5, 7, 11, 13}
	printSlice(s)

	s = s[:0]
	printSlice(s)

	s = s[:4]
	printSlice(s)

	// スライスの capacity を減らすのは、先頭を指定した場合だけ？
	s = s[2:]
	printSlice(s)

	fmt.Println("---------------------------------")

	s2 := []int{10, 20, 30, 40, 50, 60}
	printSlice(s2)

	s2 = s2[3:5]
	printSlice(s2)
}

func printSlice(s []int) {
	fmt.Printf("len=%d cap=%d %v\n", len(s), cap(s), s)
}
