package main

import (
	"fmt"
	"math"
)

func manyReturns() (int, bool, string) {
	return 20, true, "twenty"
}

func floatParts(number float64) (integerPart int, fractionalPart float64) {
	wholeNumber := math.Floor(number)
	return int(wholeNumber), number - wholeNumber
}

func main() {
	//manyReturns
	myInt, myBool, mySting := manyReturns()
	fmt.Println(manyReturns())
	fmt.Println(mySting, myInt, myBool)
	//floatParts
	cans, remainder := floatParts(9.228)
	fmt.Println(cans, remainder)
}
