package calendar

import (
	"errors"
	"unicode/utf8"
)

type Date struct {
	// Делаем поля неэкспортируемыми
	day   int
	month int
	year  int
}

type Event struct {
	title string
	Date
}

// Используем указатели в получателях,
// чтобы модифицировать переданный оригинал.
// БЕЗ PASS BY VALUE
func (d *Date) SetDay(day int) error {
	if 1 < day && day > 32 {
		return errors.New("month should be between 1..31")
	}
	d.day = day
	return nil
}

func (d *Date) SetMonth(month int) error {
	if month < 1 || month > 12 {
		return errors.New("month should be between 1..12")
	}
	d.month = month
	return nil
}

func (d *Date) SetYear(year int) error {
	if year < 1 {
		return errors.New("Invailid year")
	}
	d.year = year
	return nil
}

// Нужно добавить геттеры, Инкапсуляция -> БЕЗопасность
func (d *Date) Day() int {
	return d.day
}

func (d *Date) Month() int {
	return d.month
}

func (d *Date) Year() int {
	return d.year
}

func (e *Event) Title() string {
	return e.title
}

func (e *Event) SetTitle(title string) error {
	if utf8.RuneCountInString(title) > 30 {
		return errors.New("Invailid title")
	}
	e.title = title
	return nil
}
