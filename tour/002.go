package main

// / で区切られたのはなんだろう？サブパッケージ？
import (
	"fmt"
	"math/rand"
)

func main() {
	// rand.Intn はなに？引数はシードかな？
	fmt.Println("My favorite number is", rand.Intn(10))
}