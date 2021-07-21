package main

import "fmt"

func main() {
	fmt.Println("About one third: ", 1.0/3.0)
	fmt.Printf("About one third: %0.3f\n", 1.0/3.0) //%0.кол-во_знаков_после_запятойf
	//Formatting output with Printf and Sprintf
	//fmt.Sprintf возвращает отформатированную строку
	resultString := fmt.Sprintf("About one third: %0.2f\n", 1.0/3.0) //%0.кол-во_знаков_после_запятойf
	fmt.Printf(resultString)                                         //fmt.Print(resultString) || fmt.Println(resultString)
	//Formatting verbs
	fmt.Printf("The %s cost %d cents each.\n", "gumballs", 23)
	fmt.Printf("That will be $%f please.\n", 0.23*5)
	//%f 	Floating-point number
	//%d 	Decimal number
	//%s 	String
	//%t 	Boolean
	//%v	Any value
	//%#v	Any value, formatted as it would appear in Go program code
	//%T	Type of the supplied value (int, string, etc.)
	//%%	A literal percent sign
	fmt.Printf("Дробное число: %f\n", 3.1415)
	fmt.Printf("Целое число: %d\n", 3.1415)
	fmt.Printf("Строка: %s\n", "My name is Jeff")
	fmt.Printf("Логическая штука: %b\n", false)
	fmt.Printf("Дробь, Целое, Строка, Логическая: %v %v %v %v %v\n",
		2.718, 10, "Тупа пажилой гибон", "\t", true)
	fmt.Printf("Дробь, Целое, Строка, Логическая: %#v %#v %#v %#v %#v\n",
		2.718, 10, "Тупа пажилой гибон", "\t", true)
	fmt.Printf("Типы: %T %T %T %T %T\n",
		2.718, 10, "Тупа пажилой гибон", "\t", true)
	fmt.Printf("Пажилой процент %%\n")
}
