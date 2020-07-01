package main

import (
	"errors"
	"fmt"
)

func divide(divided float64, divisor float64) (float64, error) {
	if divisor == 0.0 {
		return 0, errors.New("На ноль делить ЗАПРЕЩЕНО")
	}
	return divided / divisor, nil
}
func main() {
	quotient, err := divide(5.6, 0.0)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Printf("%0.2f\n", quotient)
	}
}
