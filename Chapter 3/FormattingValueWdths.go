package main

import "fmt"

func main() {
	fmt.Printf("%f литров нужно\n", 1.819999999999999998)
	//Table
	fmt.Printf("%12s | %s\n", "Товар", "Цена в Центах") //%12s ограничение в 12 символов
	fmt.Println("------------------------------------------")
	fmt.Printf("%12s | %2d\n", "Марки", 5)
	fmt.Printf("%12s | %2d\n", "Скобки", 15)
	fmt.Printf("%12s | %2d\n", "Лента", 13)
}
