package main

import "fmt"

func main() {
	pow := make([]int, 63)
	for i := range pow {
		pow[i] = 1<<uint(i)
	}
	// 使わないインデックスとか値は、 _ で捨てられる（名前つけてるのに使ってないとコンパイルが通らない？ GoLand だからかな？）
	for _, value := range pow {
		fmt.Printf("%d\n", value)
	}
}
