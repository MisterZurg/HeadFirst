package compare

func Larger(a int, b int) int {
	if a < b { // Ошибка обратное сравнение!
		return a
	} else {
		return b
	}
}
