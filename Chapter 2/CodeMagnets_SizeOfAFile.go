package main

import (
	"fmt"
	"log"
	"os"
)

func main() {
	//fileInfo, _ := os.Stat("my.txt")  // _ ignores error
	fileInfo, err := os.Stat("my.txt") // now it doesn't ignore error
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(fileInfo.Size())
}
