package main

import (
	"errors"
	"fmt"
	"log"
)

func main() {
	err := errors.New("Высота не может быть отрицательной, " +
		"ВИЧ отрицательной")
	//fmt.Println(err.Error())
	fmt.Println(err)
	log.Fatal(err)
}
