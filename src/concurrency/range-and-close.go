package main

import "fmt"

func fibonacci(n int, c chan int) {
	x, y := 0, 1
	for i := 0; i < n; i++ {
		c <- x
		x, y = y, x+y
	}
	close(c) //これ以上送信する値がないため、チャネルをクローズする
}

func main() {
	c := make(chan int, 10)
	go fibonacci(cap(c), c) //cap(c)が10なので、10回繰り返される
	for i := range c { // チャネルが閉じられるまで値を繰り返し受信し続ける
		fmt.Println(i)
	}
}
