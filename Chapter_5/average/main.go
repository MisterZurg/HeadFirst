package main

import (
	"fmt"
	// Update 2
	"HeadFirst/Chapter_5/datafile"
	"log"
)

func main() {
	/*
		numbers := [3]float64{71.8, 56.2, 89.5}
		var sum float64 = 0
		for _, number := range numbers {
			sum += number
		}
	*/
	// fmt.Println(sum)
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
