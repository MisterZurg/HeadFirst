package main

import (
	"log"
	"net/http"
)

/*
func viewHandler(writer http.ResponseWriter, request *http.Request) {
	message := []byte("Hello, web!")
	// Метод Write - возвращает количество успешно записанных байт,
	// и ошибку если она была обнаружена.
	// Эскейпим кол-во, ничего полезного с ним не сделать.
	_, err := writer.Write(message)
	if err != nil {
		log.Fatal(err)
	}
}
*/

// Обновление программы для поддержки трех языков.
func write(writer http.ResponseWriter, message string) {
	_, err := writer.Write([]byte(message))
	if err != nil {
		log.Fatal(err)
	}
}

func englishHandler(writer http.ResponseWriter, request *http.Request) {
	write(writer, "Hello, web!")
}

func frenchHandler(writer http.ResponseWriter, request *http.Request) {
	write(writer, "Salut web!")
}

func hindiHandler(writer http.ResponseWriter, request *http.Request) {
	write(writer, "Namaste, web!")
}

func main() {
	// http.HandleFunc("/hello", viewHandler)
	http.HandleFunc("/hello", englishHandler)
	http.HandleFunc("/salut", frenchHandler)
	http.HandleFunc("/namaste", hindiHandler)
	err := http.ListenAndServe("localhost:8080", nil)
	log.Fatal(err)
}
