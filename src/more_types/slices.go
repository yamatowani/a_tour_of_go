package main

import "fmt"

func main() {
	// 型[]Tは型Tのスライスを表す
	primes := [6]int{2, 3, 5, 7, 11, 13}

	var s []int = primes[1:4]
	fmt.Println(s)
}
