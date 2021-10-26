package main

import "fmt"

// 「呼び出すたびに 1 ずつ大きくなる値を返す関数」を返す関数
func adder() func(int) int {
	// この sum は adder() が呼び出されるたびにキャプチャリングされる（ので、 adder() で作られた関数ごとに別々の値となる）
	sum := 0

	// この関数は「変数 sum に bind されている」と言える
	return func(x int) int {
		sum += x
		return sum
	}
}

func main() {
	// adder() を 2 つ呼び出してるので、同じ処理内容の別の関数が pos と neg に代入される
	pos, neg := adder(), adder()
	for i := 0; i < 10; i++ {
		fmt.Println(
			pos(i),
			neg(-2*i),
		)
	}
}
