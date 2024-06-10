package particle

import (
	"math/rand"
)

func sandCanMoveTo(c *Canvas, pos Pos) bool {
	return 	pos.x >= 0 && pos.x < GridWidth &&
			pos.y >= 0 && pos.y < GridHeight &&
			!ContainsPos(c.outlines, pos) &&
			!ContainsPos(c.sand, pos)
}

func handleSandInput(c *Canvas) {
	// Insert sand between prev position and current
	steps, start, stepX, stepY := getSteps(*c.input)
	for step := 0; step < steps; step++ {
		pos := getNextStep(step, start, stepX, stepY)
		if !ContainsPos(c.outlines, pos) && !ContainsPos(c.sand, pos) {
			c.sand = append(c.sand, Pos{ x: pos.x, y: pos.y, color: getRandomColor(sandColors) })
		}
	}
}

func processSand(c *Canvas) {
	removal := []Pos{}
	for i, sand := range c.sand {
		down := Pos{ x: sand.x, y: sand.y + 1 }
		if c.mode == canvasBottomless && down.y >= GridHeight {
			removal = append(removal, sand)
			continue
		}
		canGoDown := sandCanMoveTo(c, down)
		if canGoDown {
			c.sand[i].y = down.y
			continue
		}

		// Sand can only stack sandMaxStack tall. If taller, topple to a random or only available side.
		left := Pos{ x: sand.x - 1, y: sand.y }
		right := Pos{ x: sand.x + 1, y: sand.y }
		canGoLeft := sandCanMoveTo(c, left)
		canGoRight := sandCanMoveTo(c, right)
		if !canGoLeft && !canGoRight {
			continue
		}

		// Check if stacked higher than max stack
		stackedToMax := false
		for j := 1; j <= sandMaxStack; j++ {
			stackedToMax = ContainsPos(c.sand, Pos{ x: sand.x, y: sand.y + j })
			if !stackedToMax {
				break
			}
		}
		if !stackedToMax {
			continue
		}

		if canGoLeft {
			for j := 1; j <= sandMaxStack; j++ {
				canGoLeft = sandCanMoveTo(c, Pos{ x: left.x, y: left.y + j })
				if !canGoLeft {
					break
				}
			}
		}

		if canGoRight {
			for j := 1; j <= sandMaxStack; j++ {
				canGoRight = sandCanMoveTo(c, Pos{ x: right.x, y: right.y + j })
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
		c.sand = RemovePos(c.sand, sand)
	}
}
