package rendering

type Vector2 struct {
	X int
	Y int
}

type Vector3 struct {
	X int
	Y int
	Z int
}

type Cell struct {
	BgColor Vector3
	FgColor Vector3
	Glyph   rune
}

type Canvas struct {
	Cells [][]Cell
}
