package main

import (
	"HeadFirst/Chapter_8/PoolPuzzle/geo"
	"fmt"
)

func main() {
	location := geo.Coordinates{
		Latitude:  37.42,
		Longitude: -122.08,
	}
	fmt.Println("Latitude:", location.Latitude)
	fmt.Println("Longitude:", location.Longitude)

	// Part 2
	location2 := geo.Landmark{}
	location2.Name = "The Googleplex"
	location2.Latitude = 37.42
	location2.Longitude = -122.08
	fmt.Println(location2)
}
