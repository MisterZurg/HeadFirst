package prose

import (
	"fmt"
	"testing"
)

// Имя ф-ии должно начинаться с Test
func TestTwoElements(t *testing.T) { // передается указатель на testing.T
	// Вызываем метод для  testing.T, который не должен проходить
	// t.Error("no test written yet")

	list := []string{"apple", "orange"}
	want := "apple and orange"
	got := JoinWithCommas(list)
	if got != want {
		// t.Error("didn't match expected value")
		// Перенесно во вспомогательную функцию
		// t.Errorf("JoinWithCommas(%#v) = \"%s\", want \"%s\"", list, got, want)
		t.Error(errorString(list, got, want))
	}
}

func TestWithThreeElements(t *testing.T) {
	// Вызываем метод для  testing.T, который не должен проходить
	// t.Error("no test written yet")

	/*
		Apple bottom jeans boots with the fur
		JESSICA, did you sleep with your GOD DAMN teacher?
		what?
		DID YOU SLEEP WITH YOUR TEACHER
		mr wilson?
		MR WILSON.
		NO I DIDN'T
	*/
	list := []string{"apple", "orange", "pear"}
	want := "apple, orange, and pear"
	got := JoinWithCommas(list)
	if got != want {
		// t.Error("didn't match expected value")
		t.Error(errorString(list, got, want))
	}
}

func TestWithOneElement(t *testing.T) {
	list := []string{"apple"}
	want := "apple"
	got := JoinWithCommas(list)
	if got != want {
		t.Error(errorString(list, got, want))
	}
}

// Вспомогательная тествая функция
func errorString(list []string, got string, want string) string {
	return fmt.Sprintf("JoinWithCommas(%#v) = \"%s\", want \"%s\"", list, got, want)
}

// Табличные тесты - таблица входных данных, и ожидаемый результат
type testData struct {
	list []string
	want string
}

func TestJoinWithCommas(t *testing.T) {
	tests := []testData{
		testData{list: []string{}, want: ""},
		testData{list: []string{"apple", "orange"}, want: "apple and orange"},
		testData{list: []string{"apple", "orange", "pear"}, want: "apple, orange, and pear"},
		testData{list: []string{"apple"}, want: "apple"},
	}

	for _, test := range tests {
		got := JoinWithCommas(test.list)
		if got != test.want {
			t.Errorf("JoinWithCommas(%#v) = \"%s\", want \"%s\"", test.list, got, test.want)
		}
	}
}
