package main

import "golang.org/x/tour/pic"

func Pic(dx, dy int) [][]uint8 {
	dx * dy
}

func main() {
	pic.Show(Pic)
}
