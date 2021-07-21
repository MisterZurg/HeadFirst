package main

import (
	"HeadFirst/Chapter_14/prose"
	"fmt"
)

func main() {
	phrases := []string{"my parents", "a rodeo clown"}
	fmt.Println(prose.JoinWithCommas(phrases))

	phrases = []string{"my parents", "a rodeo clown", "a prize bull"}
	fmt.Println(prose.JoinWithCommas(phrases))

	phrases = []string{
		"Akatsuki wa tsudota",
		"Konan",
		"Itachi",
		"Kisame",
		"Sasori",
		"Deidara",
		"Kakuzu",
		"Hidan",
		"Orochimaru",
		"Zetsu",
	}
	fmt.Println(prose.JoinWithCommas(phrases))
}
