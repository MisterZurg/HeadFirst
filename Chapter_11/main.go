package main

import (
	"HeadFirst/Chapter_11/gadget"
	"fmt"
)

type Whistle string

func (w Whistle) MakeSound() {
	fmt.Println("Tweet!")
}

type Horn string

func (h Horn) MakeSound() {
	fmt.Println("Honk!")
}

type Robot string

func (r Robot) MakeSound() {
	fmt.Println("Beep Boop")
}

// Extra-Method
func (r Robot) Walk() {
	fmt.Println("Walk it, like I talk it (walk it)...")
}

type NoiseMaker interface {
	MakeSound()
}

func play(n NoiseMaker) {
	n.MakeSound()
	// n.Walk() не пренадлежит интерфейсу
}

// Исправим с помощью интерфейса
type Player interface {
	Play(string)
	Stop()
}

// func playList(device gadget.TapePlayer, songs []string){
func playList(device Player, songs []string) {
	for _, song := range songs {
		device.Play(song)
	}
	device.Stop()
}

func main() {
	// player := gadget.TapePlayer{}
	mixtape := []string{"Carly Rae Jepsen - Call Me Maybe", "DRACOVII - Guap", "KDA - MORE"}
	// playList(player, mixtape)

	// player2 := gadget.TapeRecoder{}
	// playList(player2, mixtape) <- cannot use player (type gadget.TypeRecoeder)
	// play(Horn("Android"))

	// После исправлений
	var player Player = gadget.TapePlayer{}
	playList(player, mixtape)
	player = gadget.TapeRecoder{}
	playList(player, mixtape)

	TryOut(gadget.TapeRecoder{})
	TryOut(gadget.TapePlayer{})
}

func TryOut(player Player) {
	player.Play("Test Track")
	player.Stop()
	// player.Record() Не входит в Player, но проверить то это дело как-то нужно :/
	// Для этих целей будем использовать TypeAssertion
	// recorder := player.(gadget.TapeRecoder) // паника, но используем ok notation
	recorder, ok := player.(gadget.TapeRecoder)
	if ok {
		recorder.Record()
	} else {
		fmt.Println("Player was not a TypeRecoder")
	}
}
