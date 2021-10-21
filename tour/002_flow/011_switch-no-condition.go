package main

import (
	"fmt"
	"time"
)

func main() {
	t := time.Now()

	// switch の後に条件を指定しないのも可能。 if-else をシンプルにした感じ
	// Kotlin の when に近い？
	switch {
	case t.Hour() < 12:
		fmt.Println("Good morning!")
	case t.Hour() < 17:
		fmt.Println("Good afternoon.")
	default:
		fmt.Println("Good evening.")
	}
}
