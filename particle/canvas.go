package particle

import (
	"fmt"
	"math"

	"github.com/hajimehoshi/ebiten/v2"
)

const (
	pixelSize    = 4 // enlarge pixels for 16-bit feel
	debugMenuBox = false
)

var (
	Width           int = 600
	Height          int = 800
	menuSize        int = 80
	gridWidth       int = Width / pixelSize
	gridHeight      int = Height / pixelSize
	pixeledMenuSize int = menuSize / pixelSize
	sandMaxStack    int = 2
)

type canvasMode int

const (
	canvasWithWalls canvasMode = iota
	canvasBottomless
)

type Canvas struct {
	input   *input
	mode    canvasMode
	walls   []position
	sand    []position
	water   []position
	spewers []spewer
}

func New() (*Canvas, error) {
	c := &Canvas{
		input: NewInput(),
		mode:  canvasBottomless,
	}
	if debugMenuBox {
		addDebugMenuBox(c)
	}
	return c, nil
}

func (c *Canvas) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return Width, Height
}

func (c *Canvas) Draw(screen *ebiten.Image) {
	if debugMenuBox {
		fmt.Println("----------REDRAW----------")
	}
	screen.Fill(backgroundColor)
	for _, water := range c.water {
		if debugMenuBox {
			printMenuDebug(water, "getWaterColor()")
		}
		drawPixel(screen, water)
	}
	for _, sand := range c.sand {
		if debugMenuBox {
			printMenuDebug(sand, "getSandColor()")
		}
		drawPixel(screen, sand)
	}
	for _, wall := range c.walls {
		if debugMenuBox {
			printMenuDebug(wall, "getWallColor()")
		}
		drawPixel(screen, wall)
	}
	for _, spewer := range c.spewers {
		spewer.draw(screen)
	}
	drawMenu(screen)
}

func (c *Canvas) Update() error {
	processInputs(c)
	processParticles(c)
	processSpewers(c)
	return nil
}

func drawPixel(screen *ebiten.Image, element position) {
	sqr := generateSquare(gridToPosition(element.x), gridToPosition(element.y), pixelSize)
	for _, pos := range sqr {
		screen.Set(pos.x, pos.y, element.color)
	}
}

func drawPixelOffset(screen *ebiten.Image, element position, offset int) {
	sqr := generateSquare(gridToPosition(element.x), gridToPosition(element.y), pixelSize)
	for _, pos := range sqr {
		screen.Set(pos.x+offset, pos.y, element.color)
	}
}

func processInputs(c *Canvas) {
	c.input.Update()
	if c.input.mouseState == mouseStateLeftClick {
		handleLeftClick(c)
	} else if c.input.mouseState == mouseStateRightClick {
		handleRightClick(c)
	}
}

func handleLeftClick(c *Canvas) {
	if c.input.mousePosY <= menuSize {
		handleMenuClick(c.input)
	} else {
		switch c.input.mode {
		case inputModeWall:
			handleWallInput(c)
		case inputModeSand:
			handleSandInput(c)
		case inputModeWater:
			handleWaterInput(c)
		case inputModeWaterSpewer:
			handleSpewerInput(c, spewerWaterType)
		case inputModeSandSpewer:
			handleSpewerInput(c, spewerSandType)
		}
	}
}

func handleRightClick(c *Canvas) {
	handleRemoveWall(c)
	handleRemoveSpewer(c)
}

func getNextStep(step int, start position, stepX, stepY float64) position {
	return position{
		x: int(math.Round(float64(start.x) + float64(step)*stepX)),
		y: int(math.Round(float64(start.y) + float64(step)*stepY)),
	}
}

func getSteps(input input) (steps int, start position, stepX, stepY float64) {
	start = position{
		x: input.mousePrevGridPosX,
		y: input.mousePrevGridPosY,
	}

	steps = max(
		abs(input.mouseGridPosX-input.mousePrevGridPosX),
		abs(input.mouseGridPosY-input.mousePrevGridPosY),
		1, // Ensure will step at least once
	)
	stepX = (float64(input.mouseGridPosX) - float64(start.x)) / float64(steps)
	stepY = (float64(input.mouseGridPosY) - float64(start.y)) / float64(steps)
	return steps, start, stepX, stepY
}

func processParticles(c *Canvas) {
	processSand(c)
	processWater(c)
}
