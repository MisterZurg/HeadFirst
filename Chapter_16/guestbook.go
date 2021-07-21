package main

import (
	"bufio"
	"fmt"
	"html/template"
	"log"
	"net/http"
	"os"
)

type Guestbook struct {
	SignatureCount int
	Signatures     []string
}

func check(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func viewHandler(writer http.ResponseWriter, request *http.Request) {
	/*
		placeholder := []byte("signature list goes here")
		_, err := writer.Write(placeholder)
		check(err)
	*/
	signatures := getStrings("Chapter_16/signatures.txt")
	fmt.Printf("%#v\n", signatures)

	html, err := template.ParseFiles("Chapter_16/view.html")
	check(err)
	guestbook := Guestbook{
		SignatureCount: len(signatures),
		Signatures:     signatures,
	}
	err = html.Execute(writer, guestbook)
	check(err)
}

func newHandler(writer http.ResponseWriter, request *http.Request) {
	html, err := template.ParseFiles("Chapter_16/new.html")
	check(err)
	err = html.Execute(writer, nil)
	check(err)
}

func createHandler(writer http.ResponseWriter, request *http.Request) {
	signature := request.FormValue("signature")
	options := os.O_WRONLY | os.O_APPEND | os.O_CREATE
	file, err := os.OpenFile("Chapter_16/signatures.txt", options, os.FileMode(0600))
	check(err)
	// _, err := writer.Write([]byte(signature))
	// check(err)
	_, err = fmt.Fprintln(file, signature) // Записывает текст в новой строке файла
	check(err)
	err = file.Close()
	check(err)
	// Редирект на главную
	http.Redirect(writer, request, "/guestbook", http.StatusFound)
}

func getStrings(fileName string) []string {
	var lines []string
	file, err := os.Open(fileName)
	// Если файл не существует
	if os.IsNotExist(err) {
		return nil
	}
	check(err)
	defer file.Close()
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	check(scanner.Err())
	return lines
}

func main() {
	http.HandleFunc("/guestbook", viewHandler)
	http.HandleFunc("/guestbook/new", newHandler)
	http.HandleFunc("/guestbook/create", createHandler)
	err := http.ListenAndServe("localhost:8080", nil)
	// Возвращаемое значение http.ListenAndServe() не передается в check(),
	// потому-что оно всегда != nil поэтому мы просто вызоваем
	log.Print(err)
}
