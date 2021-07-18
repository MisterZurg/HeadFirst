package main

import (
	"HeadFirst/Chapter_10/NotReallyPoolPuzzle/geo"
	"fmt"
	"log"
)

func main() {
	coordinates := geo.Coordinates{}
	err := coordinates.SetLatitude(37.42)
	if err != nil {
		log.Fatal(err)
	}

	err = coordinates.SetLongitude(-1137.42)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(coordinates.Latitude())
	fmt.Println(coordinates.Longitude())
}
