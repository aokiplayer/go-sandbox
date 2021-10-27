package main

import "fmt"

func main() {
	m := make(map[string]int)

	m["Answer"] = 42
	fmt.Println("The value:", m["Answer"])

	m["Answer"] = 48
	fmt.Println("The value:", m["Answer"])

	// value が int だから、 delete してキーがなくなったらデフォルトの 0 になった？ nil ではないのか
	delete(m, "Answer")
	fmt.Println("The value:", m["Answer"])

	// マップ[キー] で返される 2 つ目は、そのキーが存在するかどうか
	v, ok := m["Answer"]
	fmt.Println("The value:", v, "Present?", ok)
}
