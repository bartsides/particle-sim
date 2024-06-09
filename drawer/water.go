package drawer

import (
	"math/rand"
)

func waterCanMoveTo(c *Canvas, pos Pos) bool {
	return 	pos.x >= 0 && pos.x < (Width/pixelSize) &&
			pos.y >= 0 && pos.y < (Height/pixelSize) &&
			!ContainsPos(c.outlines, pos) &&
			!ContainsPos(c.water, pos) &&
			!ContainsPos(c.water, pos)
}

func processWater(c *Canvas) {
	for i, water := range c.water {
		// TODO: Check if water and sand share space and move up if true
		down := Pos{ x: water.x, y: water.y + 1 }
		canGoDown := waterCanMoveTo(c, down)
		if canGoDown {
			c.water[i].y = down.y
			continue
		}

		left := Pos{ x: water.x - 1, y: water.y }
		right := Pos{ x: water.x + 1, y: water.y }
		canGoLeft := waterCanMoveTo(c, left)
		canGoRight := waterCanMoveTo(c, right)
		if !canGoLeft && !canGoRight {
			continue
		}

		// If not on top of water, stop moving
		if !ContainsPos(c.water, down) {
			continue;
		}

		var fallingLeft bool
		if canGoLeft && !canGoRight {
			fallingLeft = true
		} else if !canGoLeft && canGoRight {
			fallingLeft = false
		} else {
			fallingLeft = rand.Intn(2) == 1
		}

		if fallingLeft {
			c.water[i].x = left.x
		} else {
			c.water[i].x = right.x
		}
	}
}