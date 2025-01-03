package main

import (
	"fmt"
	"math"
)

type Vertex struct {
	X, Y float64
}

// 変数もポインタもレシーバとして取る事ができる
func (v Vertex) Abs() float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

// ポインタは渡せない, ポインタの場合は*pと、値を取得して渡さないといけない
func AbsFunc(v Vertex) float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

func main() {
	v := Vertex{3, 4}
	fmt.Println(v.Abs())
	fmt.Println(AbsFunc(v))

	p := &Vertex{4, 3}
	fmt.Println(p.Abs()) // (*p).Abs()として解釈される
	fmt.Println(AbsFunc(*p))
}
