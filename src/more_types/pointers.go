package main

import "fmt"

func main() {
	i, j := 42, 2701

	p := &i // 変数iのポイント(メモリアドレス)を&演算子で引き出す
	fmt.Println(*p) // *演算子で、ポインタの指す先の変数、メモリアドレスに入っている変数を取り出す
	*p = 21 // iが入っているポインタpの場所に21を代入, 破壊的変更になる
	fmt.Println(i)

	p = &j
	*p = *p / 37 // jをポインタを通じて計算
	fmt.Println(j)
}
