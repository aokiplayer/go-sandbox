package main

import (
	"fmt"
	"runtime"
)

func main() {
	// runtime.GOOS は、プラットフォームの名前
	fmt.Println(runtime.GOOS)

	fmt.Print("Go runs on ")

	// break は不要。 C の break 忘れと同じにするには fallthrough を記述
	// Swift と異なり、全ケース網羅してない場合に default はなくても OK
	switch os := runtime.GOOS; os {
	case "darwin":
		fmt.Println("macOS.")
	case "linux":
		fmt.Println("Linux.")
	default:
		fmt.Printf("%s.\n",os)
	}
}

