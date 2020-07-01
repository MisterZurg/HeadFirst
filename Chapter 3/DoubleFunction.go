package main

import "fmt"

func main() {
	num1, num2 := 2, 2
	doubleNoPointer(num1)
	doubleWithPointer(&num2)
	fmt.Printf("NoPointer %d\n", num1)
	fmt.Printf("WithPointer %d\n", num2)
}
func doubleNoPointer(number int) {
	number *= 2
}
func doubleWithPointer(number *int) {
	*number *= 2
}
