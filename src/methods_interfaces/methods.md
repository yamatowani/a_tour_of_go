## レシーバ(receiver)とは？
- Goにおける、「メソッドをどの型に関連付けるかを定義する仕組み」
- レシーバを使うことで、特定の型に関連付けられた関数として動作する
- その型のインスタンスに対して動作するメソッドを定義できる

## コードの解説
```go
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

func main() {
	v := Vertex{3, 4}
	fmt.Println(v.Abs())
}
```
- func (v Vertex) Abs()
  - レシーバ(v Vertex)によって、AbsメソッドはVertex型に関連付けられる
  - メソッド内でv.Xやv.Yにアクセスできる
- レシーバ付きメソッドは, v.Abs()のように型のインスタンスを使って呼び出す
- 
