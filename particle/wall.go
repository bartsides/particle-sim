package particle

func handleWallInput(c *Canvas) {
	// TODO: Remove sand or water at position when adding wall?
	// Insert wall between prev position and current
	steps, start, stepX, stepY := getSteps(*c.input)
	for step := 0; step < steps; step++ {
		pos := getNextStep(step, start, stepX, stepY)
		if !containsPos(c.walls, pos) {
			c.walls = append(c.walls, position{x: pos.x, y: pos.y, color: wallColor})
		}
	}
}

func handleRemoveWall(c *Canvas) {
	pos := position{x: c.input.mouseGridPosX, y: c.input.mouseGridPosY}
	c.walls = removePos(c.walls, pos)
}
