package main

import "fmt"

type student struct {
	name  string
	grade float64
}

func printInfoStudent(s student) {
	fmt.Println("Name:", s.name)
	fmt.Printf("Grade: %0.1f\n", s.grade)
}

func main() {
	var s student
	s.name = "Alzono Cole"
	s.grade = 91.9
	printInfoStudent(s)
}
