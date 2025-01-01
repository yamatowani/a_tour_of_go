package main

import "fmt"

const (
	// 1ビットを左に100シフトすることで巨大な数値を作成
	Big = 1 << 100
	// さらに右に99シフトし、最終的には1<<1(1を1ビット左にシフトした)か2になる
	Small = Big >> 99
)

func needInt(x int) int { return x*10 + 1}
func needFloat(x float64) float64 {
	return x * 0.1
}

func main() {
	fmt.Println(needInt(Small))
	fmt.Println(needInt(Big))
	fmt.Println(needFloat(Small))
	fmt.Println(needFloat(Big))
}




