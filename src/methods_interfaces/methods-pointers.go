package main

import (
	"fmt"
	"math"
)


type Vertex struct {
	X, Y float64
}

func (v Vertex) Abs() float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

// ポインタレシーバを持つメソッドはレシーバが指す変数を変更できる。
// レシーバ自身を変更するため(変数レシーバではコピーを操作するため破壊的な変更にならない)、変数レシーバよりポインタレシーバの方が使われる
func (v *Vertex) Scale(f float64) {
	v.X = v.X * f
	v.Y = v.Y * f
}

func main() {
	v := Vertex{3, 4}
	v.Scale(10)
	fmt.Println(v.Abs())
}
