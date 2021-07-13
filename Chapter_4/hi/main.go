package main

import (
	"HeadFirst/Chapter_4/greeting"
	"HeadFirst/Chapter_4/greeting/czech"
	"HeadFirst/Chapter_4/greeting/dansk"
	"HeadFirst/Chapter_4/greeting/deutsch"
	"HeadFirst/Chapter_4/greeting/korean"
	"HeadFirst/Chapter_4/greeting/russian"
)

func main() {
	greeting.Hello()
	greeting.Hi()

	korean.Annyeong()
	korean.Annyeonghaseyo()

	deutsch.Hallo()
	deutsch.GutenTag()

	dansk.Hej()

	russian.Privet()
	russian.Zdravstvuite()

	czech.Ahoj()
	czech.DobryDen()
}
