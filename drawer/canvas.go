package drawer

import (
	"math"
	"math/rand"

	"github.com/hajimehoshi/ebiten/v2"
)

var Width int = 600
var Height int = 400
var sandMaxStack int = 2	// Must be 2 or greater

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
	for _, sand := range c.sand {
		drawPixel(screen, sand)
	}
	for _, water := range c.water {
		drawPixel(screen, water)
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
				if !Contains(c.outlines, pos) {
					c.outlines = append(c.outlines, Pos{ x: pos.x, y: pos.y, color: outlineColor })
				}
			}
		} else if c.input.mode == inputModeSand {
			// Insert sand between prev position and current
			steps, start, stepX, stepY := getSteps(*c.input)
			for step := 0; step < steps; step++ {
				pos := getNextStep(step, start, stepX, stepY)
				if !Contains(c.outlines, pos) && !Contains(c.sand, pos) {
					c.sand = append(c.sand, Pos{ x: pos.x, y: pos.y, color: sandColor })
				}
			}
		} else if c.input.mode == inputModeWater {
			// Insert water between prev position and current
			steps, start, stepX, stepY := getSteps(*c.input)
			for step := 0; step < steps; step++ {
				pos := getNextStep(step, start, stepX, stepY)
				if !Contains(c.outlines, pos) && !Contains(c.water, pos) {
					c.water = append(c.water, Pos{ x: pos.x, y: pos.y, color: waterColor })
				}
			}
		}
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
	
	// TODO: Sand can push up water. Can't stack any high without trying to topple over either side
}

func processSand(c *Canvas) {
	for i, sand := range c.sand {
		down := Pos{ x: sand.x, y: sand.y + 1 }
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
			stackedToMax = Contains(c.sand, Pos{ x: sand.x, y: sand.y + j })
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
}

func sandCanMoveTo(c *Canvas, pos Pos) bool {
	// x, y within bounds
	// no outline or sand there already
	return 	pos.x >= 0 && pos.x < (Width/pixelSize) &&
			pos.y >= 0 && pos.y < (Height/pixelSize) &&
			!Contains(c.outlines, pos) &&
			!Contains(c.sand, pos)
}
