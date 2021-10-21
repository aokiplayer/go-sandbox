package main

import "fmt"

func main() {
	var i, j int = 1, 2
	// := で代入すると、暗黙的な変数宣言（var）となる。ただし、関数内でのみ可能
	k := 3
	c, python, java := true, false, "no!"
	fmt.Println(i, j, k, c, python, java)
}
