package main

import "fmt"

var pow = []int{1, 2, 4, 8, 16, 32, 64, 128}

func main() {
	// for で利用する range は、 slice や map をひとつずつ反復処理するために利用
	// range スライス とすると、「要素のインデックス」と「要素のコピー」の 2 つが取れる
	for i, v := range pow {
		fmt.Printf("2**%d = %d\n", i, v)
	}
}
