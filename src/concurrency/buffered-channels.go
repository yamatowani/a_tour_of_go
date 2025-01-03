package main

import "fmt"

func main() {
	ch := make(chan int, 2) // 三つ目の変数はバッファの長さを与える
	ch <- 1 // chに1を送信
	ch <- 2 // chに2を送信
	fmt.Println(<-ch) //受信
	fmt.Println(<-ch)
}
