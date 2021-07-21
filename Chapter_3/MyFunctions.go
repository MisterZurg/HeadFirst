package main

import "fmt"

func sayHi() {
	fmt.Println("Говорю Прiвiт!")
}
func repeatLine(line string, times int) {
	for i := 0; i < times; i++ {
		fmt.Println(line)
	}
}
func main() {
	sayHi()
	repeatLine("Здравствуйте", 3)
}
