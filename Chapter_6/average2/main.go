package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
)

func main() {
	// fmt.Println(os.Args) // Здесь содержится имя исполняемого файла
	// fmt.Println(os.Args[1:]) // Вывод аргументов без имени исполняемого файла
	arguments := (os.Args[1:])
	// var sum float64 = 0
	var numbers []float64 // Поробуем в variadic function
	for _, argument := range arguments {
		number, err := strconv.ParseFloat(argument, 64)
		if err != nil {
			log.Fatal(err)
		}
		//sum += number
		numbers = append(numbers, number)
	}
	// sampleCount := float64(len(arguments))
	// fmt.Printf("Average: %0.2f\n", sum/sampleCount)

	// Cannot use 'numbers' (type []float64) as type float64
	// fmt.Printf("Average: %0.2f\n", average(numbers))
	fmt.Printf("Average: %0.2f\n", average(numbers...))
}

// Update 2
func average(numbers ...float64) float64 {
	var sum float64 = 0
	for _, number := range numbers {
		sum += number
	}
	return sum / float64(len(numbers))
}
