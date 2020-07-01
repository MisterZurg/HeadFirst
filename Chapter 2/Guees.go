package main

import (
	"bufio"
	"fmt"
	"log"
	"math/rand"
	"os"
	"strconv"
	"strings"
	"time"
)

func main() {
	seconds := time.Now().Unix() //Получаем текущие дату и время как число
	rand.Seed(seconds)
	target := rand.Intn(100) + 1 //Без +1 дегенерация от {0..99}
	fmt.Println("I've chosen a random number between 1 and 100.")
	fmt.Println("Can you guess it?")
	//fmt.Println(target)
	success := false
	reader := bufio.NewReader(os.Stdin)
	for guesses := 0; guesses < 10; guesses++ {
		fmt.Println("You have", 10-guesses, "guesses left.")

		fmt.Print("Make a guess: ")
		input, err := reader.ReadString('\n')
		if err != nil {
			log.Fatal(err)
		}
		input = strings.TrimSpace(input)  //Remove the new line
		guess, err := strconv.Atoi(input) //Conver 2 int
		if err != nil {
			log.Fatal(err)
		}
		if guess < target {
			fmt.Println("Opps. Ur guess is LOW")
		} else if guess > target {
			fmt.Println("Opps. Ur guess is HIGH")
		} else {
			success = true
			fmt.Println("Here in my garage, just bought this new Lamborghini here.\n" +
				"It’s fun to drive up here in the Hollywood hills.\n" +
				"But you know what I like more than materialistic things? Knowledge.\n" +
				"In fact, I’m a lot more proud of these seven new bookshelves\n" +
				"that I had to get installed to hold two thousand new books that I bought.\n " +
				"It’s like the billionaire Warren Buffett says,\n" +
				"“the more you learn, the more you earn.”. ")
			break
		}
	}
	if !success {
		fmt.Println("Sorry Link the target was", target)
	}
}
