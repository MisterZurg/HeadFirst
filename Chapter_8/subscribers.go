package main

import (
	"HeadFirst/Chapter_8/magazine"
	"fmt"
)

func main() {
	var subscriber1 magazine.Subscriber
	subscriber1.Name = "Jhon Wick"
	fmt.Println(subscriber1.Name)
	var subscriber2 magazine.Subscriber
	subscriber1.Name = "Monkey D Luffy"
	fmt.Println(subscriber2.Name)

	subscriber3 := defaultSubscriber("Bell Cranel")
	printInfo(subscriber3)

	subscriber4 := magazine.Subscriber{Name: "Aman Singh", Rate: 4.99, Active: true}
	printInfo(subscriber4)

	var employee magazine.Employee
	employee.Name = "Зубенко Михаил Петрович"
	employee.Salary = 9000 // per day

	/*
		var address magazine.Address
		address.Street = "123 Oak St"
		address.City = "Omaha"
		address.State = "NE"
		address.PostalCode = "68111"
		fmt.Println(address)
	*/

	address := magazine.Address{
		Street:     "123 Oak St",
		City:       "Omaha",
		State:      "NE",
		PostalCode: "68111",
	}

	fmt.Println("Before:", subscriber4.HomeAddress)
	subscriber4.HomeAddress = address
	fmt.Println("After:", subscriber4.HomeAddress)

}

func defaultSubscriber(name string) magazine.Subscriber {
	var s magazine.Subscriber
	s.Name = name
	s.Rate = 5.99
	s.Active = true
	return s
}

func printInfo(s magazine.Subscriber) {
	fmt.Println("Name:", s.Name)
	fmt.Println("Monthly rate:", s.Rate)
	fmt.Println("Active?", s.Active)
}

func applyDiscount(s *magazine.Subscriber) { // Получает указатель на структуру, а не саму структуру.
	// (*s).rate = 4.99
	s.Rate = 4.99
}
