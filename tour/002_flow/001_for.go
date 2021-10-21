package main

import "fmt"

func main() {
	sum:=0
	// 基本は C のと変わらない。ブロックの {} は必須
	for i := 0; i < 10; i++ {
		sum+=i
	}
	fmt.Println(sum)
}

