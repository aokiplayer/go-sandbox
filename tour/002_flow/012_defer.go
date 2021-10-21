package main

import "fmt"

func main() {
	// defer は、外側の関数を return するまで遅延される
	// Swift と同じにブロックレベルでも使えるのかな？
	defer fmt.Println("world.")

	fmt.Println("hello")
}
