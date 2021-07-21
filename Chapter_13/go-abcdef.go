package main

import "fmt"

func abc(channel chan string) {
	channel <- "a"
	channel <- "b"
	channel <- "c"
}

func def(channel chan string) {
	channel <- "d"
	channel <- "e"
	channel <- "f"
}

func main() {
	// Создадим 2 кАнала
	channel1 := make(chan string)
	channel2 := make(chan string)

	go abc(channel1)
	go def(channel2)

	// Получаем и выводим значения из кАналов по порядку
	fmt.Print(<-channel1)
	fmt.Print(<-channel2)
	fmt.Print(<-channel1)
	fmt.Print(<-channel2)
	fmt.Print(<-channel1)
	fmt.Print(<-channel2)
	fmt.Println()
}
