package main

import "fmt"

type Liters float64
type Gallons float64

type Milliliters float64

func main() {
	var carFuel Gallons
	var busFuel Liters
	carFuel = Gallons(10.0)
	busFuel = Liters(240.0)
	fmt.Println(carFuel, busFuel)

	carFuel = Gallons(Liters(40.0)) // 40 литров НЕ РАВНЫ 40 галлонам!
	busFuel = Liters(Gallons(63.0)) // 63 галлона НЕ РАВНЫ 63 литрам!
	fmt.Printf("Gallons: %0.1f Liters: %0.1f\n", carFuel, busFuel)

	/*
		Быстрый поиск в интернете показывает,
		что один литр равен приблизительно 0,264 галлона,
		а один галлон равен приблизительно 3,785 литра
	*/
	carFuel = Gallons(Liters(40.0) * 0.264)
	busFuel = Liters(Gallons(63.0) * 3.785)
	fmt.Printf("Gallons: %0.1f Liters: %0.1f\n", carFuel, busFuel)
}

/* // Можно создать ф-ии конвертации см. ToLiters и ToGallons
func GallonsToLiters(g Gallons) Liters {
	return Liters(g * 3.785)
}

func LitersToGallons(l Liters) Gallons {
	return Gallons(l * 0.264)
}

func GallonsToMilliters(g Gallons) Milliliters {
	return Milliliters(g * 3765.42)
}

func MilliterstoGallons(m Milliliters) Gallons {
	return Gallons(m * 0.000264)
}
*/

// А можно ф-ии получателей причем как-бы юзаем одинаковое название
// ToGallons ибо одно для Литров другое для Миллилитров
func (l Liters) ToGallons() Gallons {
	return Gallons(l * 0.264)
}

func (m Milliliters) ToGallons() Gallons {
	return Gallons(m * 0.000264)
}

func (g Gallons) ToLiters() Liters {
	return Liters(g * 3.785)
}

func (g Gallons) ToMilliliters() Milliliters {
	return Milliliters(g * 3765.42)
}

// Упражнение
func (l Liters) ToMilliliters() Milliliters {
	return Milliliters(l * 1000)
}

func (m Milliliters) ToLiters() Liters {
	return Liters(m / 1000)
}
