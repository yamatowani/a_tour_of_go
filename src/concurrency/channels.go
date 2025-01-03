package main

import "fmt"

func sum(s []int, c chan int) {
	sum := 0
	for _, v := range s {
		sum += v
	}
	c <- sum // 変数sumをチャンネルcに送信
}

func main() {
	s := []int{7, 2, 8, -9, 4, 0}

	c := make(chan int)
	// 半分ずつ分けて並行して処理する
	go sum(s[:len(s)/2], c)
	go sum(s[len(s)/2:], c)
	x, y := <-c, <-c // チャンネルcから受信した変数をx, yに割り当てる
	fmt.Println(x, y, x+y)
}
