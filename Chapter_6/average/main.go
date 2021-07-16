package main

import (
	"fmt"
	// Update 3       \/
	"HeadFirst/Chapter_6/datafile"
	"log"
)

func main() {

	numbers, err := datafile.GetFloats("Chapter_5/average/data.txt")
	if err != nil {
		log.Fatal(err)
	}
	var sum float64 = 0
	for _, number := range numbers {
		sum += number
	}
	// Вычисление среднего значения
	sampleCount := float64(len(numbers))
	fmt.Printf("Average %0.2f\n", sum/sampleCount)
}
