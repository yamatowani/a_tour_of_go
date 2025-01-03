package main

import "fmt"

func main() {
	var i interface{} = "hello"
	s := i.(string) //インターフェースの値iが型stringを保持し、変数sに代入する。iがstring型の変数を保持していないとエラーになる
	fmt.Println(s)

	s, ok := i.(string)
	fmt.Println(s, ok)

	f, ok := i.(float64) // 0, falseのゼロ値になる
	fmt.Println(f, ok)

	f = i.(float64)
	fmt.Println(f)
}
