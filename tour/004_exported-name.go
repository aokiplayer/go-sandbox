package main

import (
	"fmt"
	"math"
)

func main() {
	// fmt パッケージの Println 関数を実行。先頭大文字なのが export されたもの
	// 先頭小文字のは、外部パッケージからアクセス不可
	// たとえば、 math.intSize も定義されてるけどアクセスできない
	fmt.Println(math.Pi)
}
