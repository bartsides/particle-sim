package drawer

import "github.com/hajimehoshi/ebiten/v2"

type spewerType int
const (
	spewerSandType spewerType = iota
	spewerWaterType
)

type spewer struct {
	pos			Pos
	spewerType	spewerType
}

func (spewer *spewer) draw(screen *ebiten.Image) {
	sqr := generateSquare(ToPos(spewer.pos.x), ToPos(spewer.pos.y), pixelSize*4)
	for _, pos := range sqr {
		screen.Set(pos.x, pos.y, spewer.pos.color)
	}
}

func (spewer *spewer) update(c *Canvas) {
	switch spewer.spewerType {
	case spewerWaterType:
		if waterCanMoveTo(c, spewer.pos) {
			c.water = append(c.water, Pos{ 
				x: spewer.pos.x, 
				y: spewer.pos.y, 
				color: getRandomColor(waterColors),
			})
		}
	case spewerSandType:
		if sandCanMoveTo(c, spewer.pos) {
			c.sand = append(c.sand, Pos{
				x: spewer.pos.x,
				y: spewer.pos.y,
				color: getRandomColor(sandColors),
			})
		}
	}
}

func handleSpewerInput(c *Canvas, spewerType spewerType) {
	spewer := spewer{
		pos: Pos{x: c.input.mouseGridPosX, y: c.input.mouseGridPosY, color: spewerColor},
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