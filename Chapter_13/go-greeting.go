package main

import "fmt"

func greeting(myChannel chan string) {
	myChannel <- "hi" // Значение отправляется по кАналу
}

func main() {
	// Создаем новый канал
	myChannel := make(chan string)
	// Передаем его функции, выполняющейся в go-рутине
	go greeting(myChannel)
	// Вычитываем значение
	fmt.Println(<-myChannel)
}
