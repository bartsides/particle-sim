package drawer

import (
	"image/color"
	"math"

	"github.com/hajimehoshi/ebiten/v2"
)

var Width int = 600
var Height int = 400

var (
	backgroundColor = color.RGBA{0xfa, 0xf8, 0xef, 0xff}
	outlineColor     = color.RGBA{0xbb, 0xad, 0xa0, 0xff}
	sandColor		= color.RGBA{237, 188, 121, 255}
	waterColor		= color.RGBA{ 30, 129, 176, 255}
	drawWidth		= 3
)

type Canvas struct {
	input		*Input
	image		*ebiten.Image
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

func (c *Canvas) Draw(screen *ebiten.Image) {
	if c.image == nil {
		c.image = ebiten.NewImage(Width, Height)
	}
	screen.Fill(backgroundColor)
	for _, outline := range c.outlines {
		screen.Set(outline.x, outline.y, outline.color)
	}
	for _, sand := range c.sand {
		screen.Set(sand.x, sand.y, sand.color)
	}
	for _, water := range c.water {
		screen.Set(water.x, water.y, water.color)
	}
}

func (c *Canvas) Update() error {
	c.input.Update()
	if c.input.mouseState == mouseStateLeftClick {
		if c.input.mode == inputModeOutline {
			// TODO: Remove sand or water at position when adding outline?
			// Insert outline between prev position and current
			startX := float64(c.input.mousePrevPosX)
			startY := float64(c.input.mousePrevPosY)
			// Ensure will step at least once
			steps := max(abs(c.input.mousePosX - c.input.mousePrevPosX), abs(c.input.mousePosY - c.input.mousePrevPosY), 1)
			stepX := (float64(c.input.mousePosX) - startX) / float64(steps)
			stepY := (float64(c.input.mousePosY) - startY) / float64(steps)
			
			for step := 0; step < steps; step++ {
				x := int(math.Round(startX + float64(step)*stepX))
				y := int(math.Round(startY + float64(step)*stepY))
				rect := generateRect(x, y, drawWidth)
				for _, pos := range rect {
					if !Contains(c.outlines, pos.x, pos.y) {
						c.outlines = append(c.outlines, Pos{ x: pos.x, y: pos.y, color: outlineColor })
					}
				}
			}
		} else if c.input.mode == inputModeSand {
			// Insert sand
			rect := generateRect(c.input.mousePosX, c.input.mousePosY, drawWidth)
			for _, pos := range rect {
				if !Contains(c.outlines, pos.x, pos.y) && !Contains(c.sand, pos.x, pos.y) {
					c.sand = append(c.sand, Pos{ x: pos.x, y: pos.y, color: sandColor })
				}
			}
		} else if c.input.mode == inputModeWater {
			// Insert water
			rect := generateRect(c.input.mousePosX, c.input.mousePosY, drawWidth)
			for _, pos := range rect {
				if !Contains(c.outlines, pos.x, pos.y) && !Contains(c.water, pos.x, pos.y) {
					c.water = append(c.water, Pos{ x: pos.x, y: pos.y, color: waterColor })
				}
			}
		}
	}
	// TODO: Process particles
	return nil
}

func (c *Canvas) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return Width, Height
}

func generateRect(x, y, width int) []Pos {
	// Centered on x, y
	res := []Pos{{x: x, y: y}}

	for i := 1; i < width; i++ {
		left := x-i;
		right := x+i;
		top := y-i;
		bottom := y+i;

		res = fillLine(res, left, top, right, top)
		res = fillLine(res, right, top, right, bottom)
		res = fillLine(res, right, bottom, left, bottom)
		res = fillLine(res, left, bottom, left, top)
	}

	return res;
}

func fillLine(res []Pos, x1, y1, x2, y2 int) []Pos {
	// Only handles one direction
	if x1 == x2 {
		for y := min(y1, y2); y <= max(y1, y2); y++ {
			if (!Contains(res, x1, y)) {
				res = append(res, Pos{ x: x1, y: y })
			}
		}
	} else if y1 == y2 {
		for x := min(x1, x2); x <= max(x1, x2); x++ {
			if (!Contains(res, x, y1)) {
				res = append(res, Pos{ x: x, y: y1 })
			}
		}
	}
	
	return res
}

func abs(val int) int {
	if val < 0 {
		return -val
	}
	return val
}