package main

import (
	"HeadFirst/Chapter_7/datafile"
	"fmt"
	"log"
)

func main() {
	lines, err := datafile.GetStrings("Chapter_7/count/votes.txt")
	if err != nil {
		log.Fatal(err)
	}
	// fmt.Println(lines)
	// Сложный и плохой и не рабочий способ через slices
	/*
		var names []string
		var count []int
		for _, line := range lines {
			matched := false
			for i, name := range names{
				if name == line {
					count[i]++
					matched = true
				}
			}
			if matched == false {
				names = append(names, line)
				count = append(count, 1)
			}
		}
		for i, name := range names {
			fmt.Printf("%s: %d\n", name, count[i])
		}
	*/
	// Версия программы с использованием мапы
	counts := make(map[string]int)
	for _, line := range lines {
		counts[line]++
	}
	// fmt.Println(counts)
	// Для красивого вывода в НЕГАРАНТИРОВАННОМ порядке
	for key, value := range counts {
		fmt.Println(key, ":", value)
	}
	/* Можно попробовать написать что-то вроде
	var names []string
	for name := range counts {
		names = append(names, name)
	}
	sort.Strings(names)
	и ещё раз пробежаться в цикле
	*/

}
