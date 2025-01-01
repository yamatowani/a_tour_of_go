package main

import "fmt"

type Vertex struct {
	X int
	Y int	
}

func main() {
	v := Vertex{1, 2}
	p := &v // 構造体vのポイント
	p.X = 1e9 // ポインタから構造体の要素にアクセスできる
	fmt.Println(v)
}
