package main

import (
	"fmt"
)

func main() {
	//Part 1
	var myInt int
	//	fmt.Println(reflect.TypeOf(&myInt))
	var myFloat float64
	//	fmt.Println(reflect.TypeOf(&myFloat))
	var myBool bool
	//	fmt.Println(reflect.TypeOf(&myString))

	//Part 2
	var myIntPointer *int // = &myInt
	myIntPointer = &myInt //Сохранеем в поинтерПеременную адресс переменной :/
	fmt.Println(myIntPointer)
	var myFloatPointer *float64 = &myFloat
	fmt.Println(myFloatPointer)
	//Через короткое объявление
	myBoolPointer := &myBool
	fmt.Println(myBoolPointer)

	//Part 3 получаем значение переменной по Поинтеру(её адресу)
	myInt = 300
	myFloat = 14.88
	myBool = true
	fmt.Println(*myIntPointer)
	fmt.Println(*myFloatPointer)
	fmt.Println(*myBoolPointer)
	//The * operator can also be used to update the value at a pointer:
	*myIntPointer = 228
	fmt.Println("Updated pointer: ", *myIntPointer)
	fmt.Println("Updated variable: ", myInt)

	//Part 4.2
	var myOmegaFloatPointer *float64 = createPointer()
	fmt.Println(*myOmegaFloatPointer)
}

//Part 4.1 createPointerFunc
func createPointer() *float64 {
	var myOmegaFloat = 365.5
	return &myOmegaFloat
}
