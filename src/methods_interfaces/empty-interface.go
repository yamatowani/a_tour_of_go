package main

import "fmt"

func main() {
	// 空のインターフェースは未知の型の値を扱うコードで利用される
	var i interface{} // 空のインターフェースは任意の型の値を保持できる
	describe(i)

	i = 42
	describe(i)
	i = "hello"
	describe(i)
}

func describe(i interface{}) {
	fmt.Printf("(%v, %T)\n", i, i)
}
