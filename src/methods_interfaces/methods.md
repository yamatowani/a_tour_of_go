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
## メソッドと関数の違い
メソッドは、レシーバ引数を伴う関数のこと！！！！
## レシーバ引数
レシーバ引数は「このメソッドがどのデータ（型）についての動作なのかを示すもの」、動作の対象のこと。通常の関数と違って、どの型に対する動作なのかが明確になる。
