package main //main
import (
	"fmt"
)

var originalCount int = 10
var eatenCount int = 4

func main() {
	fmt.Println("I started with", +originalCount, "apples.")
	fmt.Println("Some jerk ate",
		eatenCount, "apples.")
	fmt.Println("There are", originalCount-eatenCount,
		"apples left.")
}
