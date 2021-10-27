package main

import "fmt"

func main() {
	sum := 1

	// 初期化や後処理は、書かなくても OK
	for ;sum < 1000; {
		sum += sum
	}

	// この形でも OK
	//for sum < 1000 {
	//	sum += sum
	//}
	fmt.Println(sum)
}
