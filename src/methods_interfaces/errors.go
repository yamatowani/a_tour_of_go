package main

import (
	"fmt"
	"time"
)

type MyError struct {
	When time.Time
	What string
}

// errorインターフェースを実装している
func (e *MyError) Error() string {
	return fmt.Sprintf("at %v, %s", e.When, e.What)
}

// 必ずMyError型のエラーを返す関数
func run() error { // 自動的にerrorインターフェースを確認し、Error()が実行される
	return &MyError{ 
		time.Now(),
		"It didn't work",
	}
}

func main() {
	if err := run(); err != nil {
		fmt.Println(err)
	}
}
