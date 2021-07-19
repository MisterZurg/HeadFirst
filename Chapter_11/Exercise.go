package main

import "fmt"

type Car string

func (c Car) Accelerate() {
	fmt.Println("Speeding up")
}

func (c Car) Brake() {
	fmt.Println("Stopping")
}

func (c Car) Steer(direction string) {
	fmt.Println("Turning", direction)
}

type Truck string

func (t Truck) Accelerate() {
	fmt.Println("Speeding up")
}

func (t Truck) Brake() {
	fmt.Println("Stopping")
}

func (t Truck) Steer(direction string) {
	fmt.Println("Turning", direction)
}

func (t Truck) LoadCargo(cargo string) {
	fmt.Println("Loading", cargo)
}

type Vehicle interface {
	Accelerate()
	Brake()
	Steer(string)
}

func main() {
	var vehicle Vehicle = Car("Toyota Supra")
	vehicle.Accelerate()
	vehicle.Steer("left")

	vehicle = Truck("CUMass 1337")
	vehicle.Brake()
	vehicle.Steer("right")

	TryVehicle(Truck("CAT"))
	TryVehicle(Car("Lambo"))
}

// PoolPuzzle Part
func TryVehicle(vehicle Vehicle) {
	vehicle.Accelerate()
	vehicle.Steer("LEFT")
	vehicle.Steer("RIGHT")
	vehicle.Brake()
	truck, ok := vehicle.(Truck)
	if ok {
		truck.LoadCargo("Test Cargo")
	}
}
