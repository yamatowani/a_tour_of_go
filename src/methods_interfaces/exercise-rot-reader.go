package main

import (
	"io"
	"os"
	"strings"
)

type rot13Reader struct {
	r io.Reader
}

func (r rot13Reader) Read(p []byte) (int, error) {
	n, err := r.r.Read(p) // 元のReaderからデータを読み込む
	for i := 0; i < n; i++ {
		p[i] = rot13(p[i]) // 読み込んだデータをROT13変換
	}
	return n, err
}

func main() {
	s := strings.NewReader("Lbh penpxrq gur pbqr!")
	r := rot13Reader{s}
	io.Copy(os.Stdout, &r)
}

func rot13(c byte) byte {
	switch {
	case ('A' <= c && c <= 'Z'):
		return (c - 'A' + 13) % 26 + 'A'
	case ('a' <= c && c <= 'z'):
		return (c - 'a' + 13) % 26 + 'a'
	default:
		return c
	}
}
