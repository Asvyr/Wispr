package rendering

import "fmt"

func NewCanvas(sizeX int, sizeY int) Canvas {
	outCells := make([][]Cell, sizeX)

	for i := range outCells {
		outCells[i] = make([]Cell, sizeY)
	}

	outCanvas := Canvas{
		Cells: outCells,
	}
	return outCanvas
}

func Draw(inCanv Canvas) {
	for y := range len(inCanv.Cells) {
		for x := range len(inCanv.Cells[0]) {
			SetCursorPos(x, y)
			fmt.Print("#")
		}
	}
}
