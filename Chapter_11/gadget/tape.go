package gadget

import "fmt"

type TapePlayer struct {
	Batteries string
}

func (t TapePlayer) Play(song string) {
	fmt.Println("Playing", song)
}

func (t TapePlayer) Stop() {
	fmt.Println("Stopped!")
}

type TapeRecoder struct {
	Microphones int
}

// Same as TapePlayer
func (t TapeRecoder) Play(song string) {
	fmt.Println("Playing:", song)
}

func (t TapeRecoder) Record() {
	fmt.Println("Recording")
}

// Same as TapePlayer
func (t TapeRecoder) Stop() {
	fmt.Println("Stopped!")
}
