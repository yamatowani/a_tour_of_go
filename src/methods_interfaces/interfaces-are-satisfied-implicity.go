package main

import "fmt"

type I interface {
	M()
}

type T struct {
	S string
}

// T型に関連付けられているメソッド
func (t T) M() {
	fmt.Println(t.S)
}

func main() {
	var i I = T{"Hello"} // T型はインターフェースIを暗黙的に実装している, 
	i.M()
}
