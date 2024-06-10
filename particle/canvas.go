package particle

import (
	"math"

	"github.com/hajimehoshi/ebiten/v2"
)

const (
	pixelSize = 4 // enlarge pixels for 16-bit feel
)

var (
	Width 		int = 600
	Height 		int = 400
	GridWidth 	int = Width / pixelSize
	GridHeight 	int = Height / pixelSize
	sandMaxStack int = 2
)

type canvasMode int
const (
	canvasWithWalls canvasMode = iota
	canvasBottomless
)

type Canvas struct {
	input		*input
	mode		canvasMode
	outlines	[]Pos // TODO: Convert outlines to use hash?
	sand		[]Pos
	water		[]Pos
	spewers		[]spewer
}

func New() (*Canvas, error) {
	c := &Canvas{
		input: NewInput(),
		mode: canvasBottomless,
	}
	return c, nil
}

func (c *Canvas) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return Width, Height
}

func (c *Canvas) Draw(screen *ebiten.Image) {
	screen.Fill(backgroundColor)
	for _, water := range c.water {
		drawPixel(screen, water)
	}
	for _, sand := range c.sand {
		drawPixel(screen, sand)
	}
	for _, outline := range c.outlines {
		drawPixel(screen, outline)
	}
	for _, spewer := range c.spewers {
		spewer.draw(screen)
	}
}

func drawPixel(screen *ebiten.Image, element Pos) {
	sqr := generateSquare(ToPos(element.x), ToPos(element.y), pixelSize)
	for _, pos := range sqr {
		screen.Set(pos.x, pos.y, element.color)
	}
}

func (c *Canvas) Update() error {
	processInputs(c)
	processParticles(c)
	processSpewers(c)
	return nil
}

func processInputs(c *Canvas) {
	c.input.Update()
	if c.input.mouseState == mouseStateLeftClick {
		switch c.input.mode {
		case inputModeOutline:
			handleOutlineInput(c)
		case inputModeSand:
			handleSandInput(c)
		case inputModeWater:
			handleWaterInput(c)
		case inputModeWaterSpewer:
			handleSpewerInput(c, spewerWaterType)
		case inputModeSandSpewer:
			handleSpewerInput(c, spewerSandType)
		}
	} else if c.input.mouseState == mouseStateRightClick {
		handleRemoveOutline(c)
		handleRemoveSpewer(c)
	}
}

func getNextStep(step int, start Pos, stepX, stepY float64) Pos {
	return Pos {
		x: int(math.Round(float64(start.x) + float64(step)*stepX)),
		y: int(math.Round(float64(start.y) + float64(step)*stepY)),
	}
}

func getSteps(input input) (steps int, start Pos, stepX, stepY float64) {
	start = Pos{
		x: input.mousePrevGridPosX,
		y: input.mousePrevGridPosY,
	}
	
	steps = max(
		abs(input.mouseGridPosX - input.mousePrevGridPosX), 
		abs(input.mouseGridPosY - input.mousePrevGridPosY),
		1,  // Ensure will step at least once
	)
	stepX = (float64(input.mouseGridPosX) - float64(start.x)) / float64(steps)
	stepY = (float64(input.mouseGridPosY) - float64(start.y)) / float64(steps)
	return steps, start, stepX, stepY
}

func processParticles(c *Canvas) {
	processSand(c)
	processWater(c)
}
