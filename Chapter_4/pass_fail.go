package main

import (
	"HeadFirst/Chapter_4/keyboard"
	"fmt"
	"log"
)

func main() {
	fmt.Println("Enter a grade: ")
	grade, err := keyboard.GetFloat()
	if err != nil {
		log.Fatal(err)
	}

	var status string
	if grade >= 60 {
		status = "passing"
	} else {
		status = "gailing"
	}
	fmt.Println("A grade of", grade, "is", status)
}
