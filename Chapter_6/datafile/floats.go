package datafile

import (
	"bufio"
	"os"
	"strconv"
)

func GetFloats(fileName string) ([]float64, error) { // Переходим на возвращение сегмента
	/*
		Переменная по умолчанию содержит nil.
		(Помните, «append» интерпретирует nil как пустой сегмент.
	*/
	var numbers []float64
	file, err := os.Open(fileName)
	if err != nil {
		// return numbers, err
		/*
			Возвращает nil вместо сегмента.
			(Переменная в этот момент в любом случае равна nil,
			но так этот факт становится еще более очевидным.)
		*/
		return nil, err
	}
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		// Теперь Строка преобразуется в float64 и присваивается временной переменной
		number, err := strconv.ParseFloat(scanner.Text(), 64)
		if err != nil {
			// return numbers, err
			return nil, err
		}
		// Новое число присоединяется к сегменту.
		numbers = append(numbers, number)
	}
	err = file.Close()
	if err != nil {
		return numbers, err
	}
	if scanner.Err() != nil {
		// return numbers, scanner.Err()
		return nil, scanner.Err()
	}
	return numbers, nil
}
