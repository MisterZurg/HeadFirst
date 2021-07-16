package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func main() {
	file, err := os.Open("Chapter_5/average/data.txt")
	if err != nil {
		log.Fatal(err)
	}

	scanner := bufio.NewScanner(file)
	for scanner.Scan() { // Читает строку из файла
		fmt.Println(scanner.Text()) // Выводит её текст
	}

	err = file.Close()
	if err != nil {
		log.Fatal(err)
	}

	if scanner.Err() != nil {
		// Если произошла ошибка при сканировании файла,
		// соощаем о ней и завершаем работу
		log.Fatal(scanner.Err())
	}
}
