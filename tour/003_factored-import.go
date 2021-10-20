package main

// これは factored import statement
// 別々の import で書いても同じだけど、 factored import statement が好まれるらしい
import (
	"fmt"
	"math"
)

func main() {
	// fmt.Printf と fmt.Print の違いは？後者も引数を複数指定できたけど？
	fmt.Printf("Now you have %g problems.\n", math.Sqrt(7))
}