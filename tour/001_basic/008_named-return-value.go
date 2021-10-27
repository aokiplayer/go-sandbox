package main

import "fmt"

// 関数の戻り値に名前を付けられる
func split(sum int) (x, y int) {
	x = sum * 4 / 9
	y = sum - x

	// naked return
	// 戻り値として返す変数の名前が決まってるので、 return だけ書いておけば良い
	return
}

func main() {
	fmt.Println(split(17))
}
