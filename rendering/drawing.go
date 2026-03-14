package rendering

import "fmt"

func SetCursorPos(x int, y int) {
	fmt.Printf("\033[%d;%dH", x, y)
}

func Clear() {
	fmt.Print("\033[H\033[2J")
}
