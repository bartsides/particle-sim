package drawer

func handleOutlineInput(c *Canvas) {
	// TODO: Remove sand or water at position when adding outline?
	// Insert outline between prev position and current
	steps, start, stepX, stepY := getSteps(*c.input)
	for step := 0; step < steps; step++ {
		pos := getNextStep(step, start, stepX, stepY)
		if !ContainsPos(c.outlines, pos) {
			c.outlines = append(c.outlines, Pos{x: pos.x, y: pos.y, color: outlineColor})
		}
	}
}

func handleRemoveOutline(c *Canvas) {
	pos := Pos{x: c.input.mouseGridPosX, y: c.input.mouseGridPosY}
	c.outlines = RemovePos(c.outlines, pos)
}