package particle

import (
	"math/rand"
)

func sandCanMoveTo(c *Canvas, pos position) bool {
	return pos.x >= 0 && pos.x < gridWidth &&
		pos.y >= 0 && pos.y < gridHeight &&
		!containsPos(c.walls, pos) &&
		!containsPos(c.sand, pos)
}

func handleSandInput(c *Canvas) {
	// Insert sand between prev position and current
	steps, start, stepX, stepY := getSteps(*c.input)
	for step := 0; step < steps; step++ {
		pos := getNextStep(step, start, stepX, stepY)
		if !containsPos(c.walls, pos) && !containsPos(c.sand, pos) {
			c.sand = append(c.sand, position{x: pos.x, y: pos.y, color: getSandColor()})
		}
	}
}

func processSand(c *Canvas) {
	removal := []position{}
	for i, sand := range c.sand {
		down := position{x: sand.x, y: sand.y + 1}
		if c.mode == canvasBottomless && down.y >= gridHeight {
			removal = append(removal, sand)
			continue
		}
		canGoDown := sandCanMoveTo(c, down)
		if canGoDown {
			c.sand[i].y = down.y
			continue
		}

		// Sand can only stack sandMaxStack tall. If taller, topple to a random or only available side.
		left := position{x: sand.x - 1, y: sand.y}
		right := position{x: sand.x + 1, y: sand.y}
		canGoLeft := sandCanMoveTo(c, left)
		canGoRight := sandCanMoveTo(c, right)
		if !canGoLeft && !canGoRight {
			continue
		}

		// Check if stacked higher than max stack
		stackedToMax := false
		for j := 1; j <= sandMaxStack; j++ {
			stackedToMax = containsPos(c.sand, position{x: sand.x, y: sand.y + j})
			if !stackedToMax {
				break
			}
		}
		if !stackedToMax {
			continue
		}

		if canGoLeft {
			for j := 1; j <= sandMaxStack; j++ {
				canGoLeft = sandCanMoveTo(c, position{x: left.x, y: left.y + j})
				if !canGoLeft {
					break
				}
			}
		}

		if canGoRight {
			for j := 1; j <= sandMaxStack; j++ {
				canGoRight = sandCanMoveTo(c, position{x: right.x, y: right.y + j})
				if !canGoRight {
					break
				}
			}
		}

		if !canGoLeft && !canGoRight {
			continue
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
			c.sand[i].x = left.x
		} else {
			c.sand[i].x = right.x
		}
	}

	for _, sand := range removal {
		c.sand = removePos(c.sand, sand)
	}
}
