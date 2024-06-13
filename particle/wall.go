package particle

func handleWallInput(c *Canvas) {
	// TODO: Remove sand or water at position when adding wall?
	// Insert wall between prev position and current
	steps, start, stepX, stepY := getSteps(*c.input)
	for step := 0; step < steps; step++ {
		pos := getNextStep(step, start, stepX, stepY)
		if !ContainsPos(c.walls, pos) {
			c.walls = append(c.walls, Pos{x: pos.x, y: pos.y, color: wallColor})
		}
	}
}

func handleRemoveWall(c *Canvas) {
	pos := Pos{x: c.input.mouseGridPosX, y: c.input.mouseGridPosY}
	c.walls = RemovePos(c.walls, pos)
}