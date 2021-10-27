package main

import "fmt"

func main() {
	fmt.Println("counting")

	// defer はスタックするので、複数指定されていた場合は後のやつから先に実行される
	// なので、ここでは 0 からカウントアップしてるけど 9 から表示される
	for i := 0; i < 10; i++ {
		defer fmt.Println(i)
	}

	fmt.Println("done")
}
