package particle

import "github.com/hajimehoshi/ebiten/v2"

type spewerType int

const (
	spewerSandType spewerType = iota
	spewerWaterType
)

type spewer struct {
	pos        position
	spewerType spewerType
}

func (spewer *spewer) draw(screen *ebiten.Image) {
	sqr := generateSquare(gridToPosition(spewer.pos.x), gridToPosition(spewer.pos.y), pixelSize*4)
	for _, pos := range sqr {
		screen.Set(pos.x, pos.y, spewer.pos.color)
	}
}

func (spewer *spewer) update(c *Canvas) {
	switch spewer.spewerType {
	case spewerWaterType:
		if waterCanMoveTo(c, spewer.pos) {
			c.water = append(c.water, position{
				x:     spewer.pos.x,
				y:     spewer.pos.y,
				color: getWaterColor(),
			})
		}
	case spewerSandType:
		if sandCanMoveTo(c, spewer.pos) {
			c.sand = append(c.sand, position{
				x:     spewer.pos.x,
				y:     spewer.pos.y,
				color: getSandColor(),
			})
		}
	}
}

func handleSpewerInput(c *Canvas, spewerType spewerType) {
	spewer := spewer{
		pos:        position{x: c.input.mouseGridPosX, y: c.input.mouseGridPosY, color: spewerColor},
		spewerType: spewerType,
	}
	if !containsSpewer(c.spewers, spewer) {
		c.spewers = append(c.spewers, spewer)
	}
}

func containsSpewer(values []spewer, spewer spewer) bool {
	for _, val := range values {
		if val.pos.x == spewer.pos.x && val.pos.y == spewer.pos.y {
			return true
		}
	}
	return false
}

func processSpewers(c *Canvas) {
	for _, spewer := range c.spewers {
		spewer.update(c)
	}
}

func handleRemoveSpewer(c *Canvas) {
	// TODO: Remove spewer
}
