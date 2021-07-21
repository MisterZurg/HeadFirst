package main

import (
	"fmt"
	"time"
)

func repeat(s string) {
	for i := 0; i < 25; i++ {
		fmt.Print(s)
	}
}

func main() {
	go repeat("x")
	go repeat("y")
	time.Sleep(time.Second)
}
