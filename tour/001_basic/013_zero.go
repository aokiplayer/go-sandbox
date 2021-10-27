package main

import "fmt"

func main() {
	// 初期値を与えないと zero value となる。型によって値は異なるけど、何もない感じ
	var (
		i int
		f float64
		b bool
		s string
	)

	fmt.Printf("%v %v %v %q\n", i, f, b, s)
}
