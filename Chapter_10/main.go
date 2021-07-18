package main

import (
	"HeadFirst/Chapter_10/calendar"
	"fmt"
	"log"
)

//type Date struct {
//	day int
//	month int
//	year int
//}

func main() {
	date := calendar.Date{}
	// date.SetDay(24)
	fmt.Println(date)

	err := date.SetYear(2021)
	if err != nil {
		log.Fatal(err)
	}

	// Все ещё можно задать нелегитимные параметры структуры Date
	// если создать ручкми и их забить.
	// с этой целью вынесем структуру в отдельный пакет и сделаем её поля не экспортируемыми
	// calendar.Date{-1, 22, 131} не удастся :!
	// теперь только сетами

	date2 := calendar.Date{}
	err2 := date2.SetDay(1)
	if err2 != nil {
		log.Fatal(err2)
	}

	event := calendar.Event{}
	err3 := event.SetYear(2021)
	if err != nil {
		log.Fatal(err3)
	}

	fmt.Println(event.Year())
}

//
//// Используем указатели в получателях,
//// чтобы модифицировать переданный оригинал.
//// БЕЗ PASS BY VALUE
//func (d *Date) SetDay(day int) error{
//	if 1 < day && day > 32 {
//		return errors.New("month should be between 1..31")
//	}
//	d.day = day
//	return nil
//}
//
//func (d *Date) SetMonth(month int) error{
//	if month < 1 || month > 12 {
//		return errors.New("month should be between 1..12")
//	}
//	d.month = month
//	return nil
//}
//
//func (d *Date) SetYear(year int) error{
//	if year < 1 {
//		return errors.New("Invailid year")
//	}
//	d.year = year
//	return nil
//}
