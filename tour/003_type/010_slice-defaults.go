package main

import "fmt"

func main() {
	s := []int{2, 3, 5, 7, 11, 13}

	s = s[1:4]
	fmt.Println(s)

	// スライスの下限値は、指定しないと 0
	s = s[:2]
	fmt.Println(s)

	// スライスの上限値は、指定しないとスライスの長さ
	s = s[1:]
	fmt.Println(s)
}
