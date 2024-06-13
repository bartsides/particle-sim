package particle

import (
	"math/rand"
)

func waterCanMoveTo(c *Canvas, pos Pos) bool {
	return 	pos.x >= 0 && pos.x < gridWidth &&
			pos.y >= 0 && pos.y < gridHeight &&
			!ContainsPos(c.walls, pos) &&
			!ContainsPos(c.sand, pos) &&
			!ContainsPos(c.water, pos)
}

func handleWaterInput(c *Canvas) {
	// Insert water between prev position and current
	steps, start, stepX, stepY := getSteps(*c.input)
	for step := 0; step < steps; step++ {
		pos := getNextStep(step, start, stepX, stepY)
		if !ContainsPos(c.walls, pos) && !ContainsPos(c.water, pos) {
			c.water = append(c.water, Pos{ x: pos.x, y: pos.y, color: getWaterColor() })
		}
	}
}

func processWater(c *Canvas) {
	removal := []Pos{}
	for i, water := range c.water {
		if ContainsPos(c.sand, water) || ContainsPos(c.walls, water) {
			// Sand or wall has displaced water
			c.water[i].y = max(0, water.y - 1)
			continue
		}

		down := Pos{ x: water.x, y: water.y + 1 }
		if c.mode == canvasBottomless && down.y >= gridHeight {
			removal = append(removal, water)
			continue
		}
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

		// Check if able to drop either left or right and nudge towards that direction
		canDropLeft := false
		canDropRight := false
		stepsLeft := 1
		stepsRight := 1

		if canGoLeft {
			canDropLeft = waterCanMoveTo(c, Pos{ x: left.x, y: left.y + 1 })
			if !canDropLeft {
				// Continue searching until left is blocked or left down is open
				for j := left.x; j >= 0; j-- {
					if !waterCanMoveTo(c, Pos{ x: j, y: left.y }) {
						// Left is blocked
						break
					}
					canDropLeft = waterCanMoveTo(c, Pos{ x: j, y: left.y + 1 })
					if canDropLeft {
						// Can drop left
						stepsLeft = j
						break
					}
				}
			}
		}

		if canGoRight {
			canDropRight = waterCanMoveTo(c, Pos{ x: right.x, y: right.y + 1 })
			if !canDropRight {
				for j := right.x; j < gridWidth; j++ {
					if !waterCanMoveTo(c, Pos{ x: j, y: right.y }) {
						// Right is blocked
						break
					}
					canDropRight = waterCanMoveTo(c, Pos{ x: j, y: right.y + 1 })
					if canDropRight {
						// Can drop right
						stepsRight = j
						break
					}
				}
			}
		}

		if !canDropLeft && !canDropRight {
			continue;
		}

		var fallingLeft bool
		if canDropLeft && !canDropRight {
			fallingLeft = true
		} else if !canDropLeft && canDropRight {
			fallingLeft = false
		} else if stepsLeft < stepsRight {
			fallingLeft = true
		} else if stepsRight > stepsLeft {
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

	for _, water := range removal {
		c.water = RemovePos(c.water, water)
	}
}