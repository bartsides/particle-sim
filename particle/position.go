package particle

import (
	"image/color"
)

type position struct {
	x     int
	y     int
	color color.Color
}

func containsPos(values []position, pos position) bool {
	for _, val := range values {
		if val.x == pos.x && val.y == pos.y {
			return true
		}
	}
	return false
}

func removePos(values []position, pos position) []position {
	for i, val := range values {
		if val.x == pos.x && val.y == pos.y {
			return append(values[:i], values[i+1:]...)
		}
	}
	return values
}

func gridToPosition(grid int) int {
	return grid * pixelSize
}

// func positionToGrid(pos int) int {
// 	return pos / pixelSize
// }

func generateSquare(x, y, size int) []position {
	// x,y is top left of rect
	res := make([]position, size*size)
	i := 0
	for x1 := 0; x1 < size; x1++ {
		for y1 := 0; y1 < size; y1++ {
			res[i] = position{x: x1 + x, y: y1 + y}
			i++
		}
	}

	return res
}
