package main

import "fmt"

func fibonacci() func() int {
	v0 := 0 // i-2 番目の数
	v1 := 1 // i-1 番目の数
	n := 0  // i 番目の数

	// フィボナッチ数を返す関数
	return func() int {
		defer func() {
			v0 = v1
			v1 = n
			n = v0 + v1
		}()

		switch n {
		case 0:
			return 0
		case 1:
			return 1
		default:
			return n
		}
	}
}

func main() {
	f := fibonacci()
	for i := 0; i < 10; i++ {
		fmt.Println(f())
	}
}
