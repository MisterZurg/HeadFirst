package geo

type Coordinates struct {
	Latitude  float64
	Longitude float64
}

// Part 2
type Landmark struct {
	Name string
	Coordinates
}
