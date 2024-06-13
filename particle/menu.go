package particle

import (
	"fmt"
	"image/color"

	"github.com/hajimehoshi/ebiten/v2"
)

const menuBorderWidth = 4

func handleMenuClick(input *input) {
	x := input.mousePosX
	if x <= menuSize {
		input.mode = inputModeWall
	} else if x > menuSize && x <= menuSize*2 {
		input.mode = inputModeSand
	} else if x > menuSize*2 && x <= menuSize*3 {
		input.mode = inputModeWater
	} else if x > menuSize*3 && x <= menuSize*4 {
		input.mode = inputModeSandSpewer
	} else if x > menuSize*4 && x <= menuSize*5 {
		input.mode = inputModeWaterSpewer
	}
}

func drawMenu(screen *ebiten.Image) {
	drawMenuBottomBorder(screen)
	drawWallButton(screen)
	drawSandButton(screen)
	drawWaterButton(screen)
	drawSandSpewerButton(screen)
	drawWaterSpewerButton(screen)
}

func drawMenuBottomBorder(screen *ebiten.Image) {
	for x := 0; x < Width; x++ {
		for y := 0; y < menuBorderWidth; y++ {
			screen.Set(x, menuSize-1-y, menuBorderColor)
		}
	}
}

func drawWallButton(screen *ebiten.Image) {
	offset := 0
	for _, point := range wallIconPoints {
		drawPixelOffset(screen, point, offset)
	}
	drawButtonOutline(screen, offset, wallColor)
}

func drawSandButton(screen *ebiten.Image) {
	offset := menuSize
	for _, point := range sandIconPoints {
		drawPixelOffset(screen, point, offset)
	}
	drawButtonOutline(screen, offset, sandColors[0])
}

func drawWaterButton(screen *ebiten.Image) {
	offset := menuSize * 2
	for _, point := range waterIconPoints {
		drawPixelOffset(screen, point, offset)
	}
	drawButtonOutline(screen, offset, waterColors[0])
}

func drawSandSpewerButton(screen *ebiten.Image) {
	offset := menuSize * 3
	drawButtonOutline(screen, offset, sandColors[1])
}

func drawWaterSpewerButton(screen *ebiten.Image) {
	offset := menuSize * 4
	drawButtonOutline(screen, offset, waterColors[1])
}

func drawButtonOutline(screen *ebiten.Image, offset int, color color.RGBA) {
	for i := 0; i < menuSize; i++ {
		x := i + offset
		for j := 0; j < menuBorderWidth; j++ {
			screen.Set(x, j, color)                   // Top
			screen.Set(x, menuSize-1-j, color)        // Bottom
			screen.Set(offset+j, i, color)            // Left
			screen.Set(menuSize-1+offset-j, i, color) // Right
		}
	}
}

func addDebugMenuBox(c *Canvas) {
	for i := 0; i < pixeledMenuSize; i++ {
		c.walls = append(c.walls,
			Pos{x: i, y: pixeledMenuSize, color: getWallColor()},                       // Top
			Pos{x: i, y: pixeledMenuSize * 2, color: getWallColor()},                   // Bottom
			Pos{x: 0, y: pixeledMenuSize + i, color: getWallColor()},                   // Left
			Pos{x: pixeledMenuSize - 1, y: pixeledMenuSize + i, color: getWallColor()}, // Right
		)
	}
}

func printMenuDebug(pos Pos, color string) {
	if pos.x > 0 && pos.x < pixeledMenuSize-1 && pos.y > pixeledMenuSize && pos.y < pixeledMenuSize*2 {
		fmt.Println("{x:", pos.x, ", y:", pos.y-pixeledMenuSize-1, ", color: ", color, "},")
	}
}
