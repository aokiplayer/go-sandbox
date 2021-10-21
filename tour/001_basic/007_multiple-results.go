package main

import "fmt"

// 関数からは複数の戻り値を返せる
// Swift のタプルや Haskell の Result みたいな感じなのかな？後者は成功と失敗だからちょっと違うかも
func swap(x, y string) (string, string) {
	return y, x
}

func main() {
	a, b := swap("hello", "world")
	// Println は可変長引数？
	fmt.Println(a, b)
}
