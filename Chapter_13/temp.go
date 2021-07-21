package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	// "time" Депрекатед к концу главы
)

/* Old
func responseSize(url string){
	fmt.Println("Getting", url)
	response, err := http.Get(url)
	if err != nil {
		log.Fatal(err)
	}
	// Отложенно особождаем ресурсы,
	// выделенные на сетевое подключение
	defer response.Body.Close()
	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		log.Fatal(err)
	}
	// Поскольку ioutil.ReadAll(response.Body) возвращает
	// []byte, нужно преобразовать его в осмысленный текст
	// fmt.Println(string(body))
	// fmt.Println(body)

	fmt.Println(len(body))
}
*/

/*
func responseSize(url string, channel chan int) {
	fmt.Println("Getting", url)
	response, err := http.Get(url)
	if err != nil {
		log.Fatal(err)
	}
	defer response.Body.Close()
	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		log.Fatal(err)
	}
	// Теперь мы передаем размер по кАналу
	channel <- len(body)
}
*/

func responseSize(url string, channel chan Page) {
	fmt.Println("Getting", url)
	response, err := http.Get(url)
	if err != nil {
		log.Fatal(err)
	}
	defer response.Body.Close()
	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		log.Fatal(err)
	}
	// Теперь мы передаем структуру Page
	channel <- Page{
		URL:  url,
		Size: len(body),
	}
}

func main() {
	/*
		// Преобразуем responseSize() для использования go-рутин
		go responseSize("https://example.com/")
		go responseSize("https://golang.org/")
		go responseSize("https://golang.org/doc")
		// time.Sleep(5 * time.Second) // Тайм Слип - далеко не идеальный
		// способ ожидания завершения go-рутин

		// go-рутины не могут использоваться с возвращаемыми значениями :/
	*/

	// Создадим кАнал для значений int
	// sizes := make(chan int)
	/*
		go responseSize("https://example.com/", sizes)
		go responseSize("https://golang.org/",sizes)
		go responseSize("https://golang.org/doc",sizes)
		fmt.Println(<-sizes)
		fmt.Println(<-sizes)
		fmt.Println(<-sizes)
	*/
	// Update 3 - Элегантное решение
	// Переносим урлы в слайс
	urls := []string{
		"https://example.com/",
		"https://golang.org/",
		"https://golang.org/doc",
	}
	// Вызываем responseSize() для каждого урла
	/*
		for _, url := range urls {
			go responseSize(url, sizes)
		}
		for i := 0; i < len(urls); i++ {
			fmt.Println(<-sizes)
		}
	*/
	pages := make(chan Page)
	for _, url := range urls {
		go responseSize(url, pages)
	}
	for i := 0; i < len(urls); i++ {
		page := <-pages
		fmt.Printf("%s: %d\n", page.URL, page.Size)
	}
}

type Page struct {
	URL  string
	Size int
}
