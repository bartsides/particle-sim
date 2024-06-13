package particle

import (
	"image/color"
)

type Pos struct {
	x     int
	y     int
	color color.Color
}

func ContainsPos(values []Pos, pos Pos) bool {
	for _, val := range values {
		if val.x == pos.x && val.y == pos.y {
			return true
		}
	}
	return false
}

func RemovePos(values []Pos, pos Pos) []Pos {
	for i, val := range values {
		if val.x == pos.x && val.y == pos.y {
			return append(values[:i], values[i+1:]...)
		}
	}
	return values
}

func ToGrid(pos int) int {
	return pos / pixelSize
}

func ToGrid64(pos float64) float64 {
	return pos / float64(pixelSize)
}

func ToPos(grid int) int {
	return grid * pixelSize
}

func ToPos64(grid float64) float64 {
	return grid * pixelSize
}

func generateSquare(x, y, size int) []Pos {
	// x,y is top left of rect
	res := make([]Pos, size*size)
	i := 0
	for x1 := 0; x1 < size; x1++ {
		for y1 := 0; y1 < size; y1++ {
			res[i] = Pos{x: x1 + x, y: y1 + y}
			i++
		}
	}

	return res
}
