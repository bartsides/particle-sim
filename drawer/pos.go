package drawer

import (
	"image/color"
)

type Pos struct {
	x 		int
	y 		int
	color 	color.Color
}

func Contains(values []Pos, x, y int) bool {
	for _, pos := range values {
		if pos.x == x && pos.y == y {
			return true
		}
	}
	return false
}