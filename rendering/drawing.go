package rendering

import "fmt"

func SetCursorPos(x int, y int) {
	fmt.Printf("\033[%d;%dH", x, y)
}
