package main

import "fmt"

const Pi = 3.14

func main() {
	// 定数は const で宣言
	// character, string, boolean, numeric 型のみ宣言可能
	// := は使えない（自動的に var になるから）
	const World = "世界"
	fmt.Println("Hello", World)
	fmt.Println("Happy", Pi, "Day")

	const Truth = true
	fmt.Println("Go rules?", Truth)
}
