package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func main() {
	fmt.Print("Enter a grade: ")
	reader := bufio.NewReader(os.Stdin)
	//Example 1
	//input1 := reader.ReadString('\n')  //<- Нехватает переданных аргументов # '\n'
	// ввод до новой строки, энтер конец ввода
	//Example 2
	//input2, err := bufio.NewReader(os.Stdin).ReadString('\n')

	//Example 3
	//input3, _ := bufio.NewReader(os.Stdin).ReadString('\n') //<- Нехватает переданных аргументов

	//Example 4
	//input4, err := reader.ReadString('\n')     		//<- err принимает nil что успешно, и нас выкидывает
	//log.Fatal(err)
	//fmt.Println(input4)  //Report the error && stops the program

	//Example 5
	input5, err := reader.ReadString('\n')
	if err != nil { //<- Если err не nil то дропаемся
		log.Fatal(err)
	}
	fmt.Println(input5)

	//My Example
	//fmt.Scan(&myInput)
	//fmt.Println(input)

	//let's Trim the newline character
	input5 = strings.Trim(input5, "\n")
	grade, err := strconv.ParseFloat(input5, 64)
	if err != nil {
		log.Fatal(err)
	}
	if grade >= 60 {
		status := "passing"
		fmt.Println(status)
	} else {
		status := "faling"
		fmt.Println(status)

	}
}
