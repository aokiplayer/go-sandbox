package main

// これのインポートが分からないので保留
import "golang.org/x/tour/wc"

func WordCount(s string) map[string]int {
	return map[string]int{"x": 1}
}

func main() {
	wc.Test(WordCount)
}
