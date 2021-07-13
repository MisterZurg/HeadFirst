// Package keyboard reads user input from the keyboard.
package keyboard

import (
	"bufio"
	"os"
	"strconv"
	"strings"
)

// GetFloat reads a floating-point number from the keyboard.
// It returns the number read and any error encountered.
func GetFloat() (float64, error) {
	reader := bufio.NewReader(os.Stdin)
	input, err := reader.ReadString('\n')
	if err != nil {
		return 0, err
	}

	input = strings.TrimSpace(input)
	number, err := strconv.ParseFloat(input, 64)
	if err != nil {
		return 0, err
	}
	return number, nil
}

// go doc packageName
// Для отображения документирующих комментариев
// для пакета go doc HeadFirst/Chapter_4/keyboard
// для функции go doc HeadFirst/Chapter_4/keyboard GetFloat

/*
При добавлении документирующих комментариев следует соблюдать ряд правил:
	• Комментарии должны состоять из полноценных предложений.
	• Комментарии для пакетов должны начинаться со слова «Package», за которым следует имя пакета:
		// Package mypackage enables widget management.
	• Комментарии для функций должны начинаться с имени функции, которую они описывают:
		// MyFunction converts widgets to gizmos.
	• В комментарии также можно включать примеры кода, которые должны снабжаться отступами.
	• Не включайте дополнительные символы для выразительности или форматирования (кроме отступов в примерах кода).
		Документирующие комментарии будут выводиться в виде простого текста и должны форматироваться соответствующим образом
*/
