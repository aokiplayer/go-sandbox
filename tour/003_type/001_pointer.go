package main

import "fmt"

func main() {
	i, j := 42, 2701

	p := &i         // i へのポインタ
	fmt.Println(*p) // ポインタ p の指すアドレスの値（つまり変数 i の値）を表示
	*p = 21         // ポインタ p の指すアドレスの位置（つまり変数 i）に値を代入
	fmt.Println(i)

	p = &j       // j へのポインタ
	*p = *p / 37 // ポインタ p の指すアドレスの値（つまり変数 j の値）を更新
	fmt.Println(j)
}
