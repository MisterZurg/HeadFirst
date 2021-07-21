package main

import "fmt"

func main() {
	var myInt int = 42
	var myIntPointer *int = &myInt
	fmt.Println(*myIntPointer)

	//main
	//	double := 2
	//	fmt.Println(double(double2))
}

//func double(number int) int {
//	return number * 2
//}
