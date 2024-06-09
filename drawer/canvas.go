package drawer

import (
	"math"

	"github.com/hajimehoshi/ebiten/v2"
)

var (
	Width int 	= 600
	Height int 	= 400
	sandMaxStack int 	= 2
)

const (
	pixelSize = 4 // enlarge pixels for 16-bit feel
)

type Canvas struct {
	input		*Input
	outlines	[]Pos // TODO: Convert outlines to use hash
	sand		[]Pos
	water		[]Pos
}

func New() (*Canvas, error) {
	c := &Canvas{
		input: NewInput(),
		outlines: []Pos{},
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
	return nil
}

func processInputs(c *Canvas) {
	c.input.Update()
	if c.input.mouseState == mouseStateLeftClick {
		if c.input.mode == inputModeOutline {
			// TODO: Remove sand or water at position when adding outline?
			// Insert outline between prev position and current
			steps, start, stepX, stepY := getSteps(*c.input)
			for step := 0; step < steps; step++ {
				pos := getNextStep(step, start, stepX, stepY)
				if !ContainsPos(c.outlines, pos) {
					c.outlines = append(c.outlines, Pos{ x: pos.x, y: pos.y, color: outlineColor })
				}
			}
		} else if c.input.mode == inputModeSand {
			// Insert sand between prev position and current
			steps, start, stepX, stepY := getSteps(*c.input)
			for step := 0; step < steps; step++ {
				pos := getNextStep(step, start, stepX, stepY)
				if !ContainsPos(c.outlines, pos) && !ContainsPos(c.sand, pos) {
					c.sand = append(c.sand, Pos{ x: pos.x, y: pos.y, color: sandColor })
				}
			}
		} else if c.input.mode == inputModeWater {
			// Insert water between prev position and current
			steps, start, stepX, stepY := getSteps(*c.input)
			for step := 0; step < steps; step++ {
				pos := getNextStep(step, start, stepX, stepY)
				if !ContainsPos(c.outlines, pos) && !ContainsPos(c.water, pos) {
					c.water = append(c.water, Pos{ x: pos.x, y: pos.y, color: waterColor })
				}
			}
		}
	} else if c.input.mouseState == mouseStateRightClick {
		pos := Pos{ x: c.input.mouseGridPosX, y: c.input.mouseGridPosY }
		c.outlines = RemovePos(c.outlines, pos)
	}
}

func getNextStep(step int, start Pos, stepX, stepY float64) Pos {
	return Pos {
		x: int(math.Round(float64(start.x) + float64(step)*stepX)),
		y: int(math.Round(float64(start.y) + float64(step)*stepY)),
	}
}

func getSteps(input Input) (steps int, start Pos, stepX, stepY float64) {
	start = Pos{
		x: input.mousePrevGridPosX,
		y: input.mousePrevGridPosY,
	}
	// Ensure will step at least once
	steps = max(
		abs(input.mouseGridPosX - input.mousePrevGridPosX), 
		abs(input.mouseGridPosY - input.mousePrevGridPosY),
		1,
	)
	stepX = (float64(input.mouseGridPosX) - float64(start.x)) / float64(steps)
	stepY = (float64(input.mouseGridPosY) - float64(start.y)) / float64(steps)
	return steps, start, stepX, stepY
}

func processParticles(c *Canvas) {
	processSand(c)
	processWater(c)
}