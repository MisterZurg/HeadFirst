package main

import "fmt"

func paintNeeded(width float64, height float64) (float64, error) { //Последний float64
	if width < 0 {
		return 0, fmt.Errorf("Ширина %0.2f - невалидна", width)
	}
	if height < 0 {
		return 0, fmt.Errorf("Высота %0.2f - невалидна", height)
	}
	area := width * height //тип возвращаемого аргумента
	//fmt.Printf("%.2f литров нужно.\n", area/10.0)
	return area / 10.0, nil
}
func main() {

	///var width , height, area float64
	//width=4.2
	//height=3.0
	//area=width*height
	//fmt.Printf("%.2f литров нужно.\n", area/10.0)
	//width=5.2
	//height=3.5
	//area=width*height
	//fmt.Printf("%.2f литров нужно.\n", area/10.0)
	//Все выше ~ вот этому
	//paintNeeded(4.2, 3.0)
	//paintNeeded(5.2, 3.5)
	//paintNeeded(5.0, 3.3)
	//Part 3
	//var amount, total float64
	//amount = paintNeeded(4.2, 3.0)
	//fmt.Printf("%0.2f литров нужно\n", amount)
	//total+=amount
	//amount = paintNeeded(5.2, 3.5)
	//fmt.Printf("%0.2f литров нужно\n", amount)
	//total+=amount
	//fmt.Printf("Всего: %0.2f литров\n",total)
	//Part 4
	amount, err := paintNeeded(4.2, -3.0)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Printf("%0.2f liters needed\n", amount)
	}

}
