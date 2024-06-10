package particle

import (
	"github.com/hajimehoshi/ebiten/v2"
)

type mouseState int
const (
	mouseStateNone mouseState = iota
	mouseStateLeftClick
	mouseStateRightClick
)

type inputMode int
const (
	inputModeOutline inputMode = iota
	inputModeSand
	inputModeWater
	inputModeWaterSpewer
	inputModeSandSpewer
)

// input represents the current key states.
type input struct {
	mouseState			mouseState
	mousePosX			int
	mousePosY			int
	mouseGridPosX		int
	mouseGridPosY		int
	mousePrevPosX 		int
	mousePrevPosY 		int
	mousePrevGridPosX 	int
	mousePrevGridPosY 	int
	mode				inputMode
}

// NewInput generates a new Input object.
func NewInput() *input {
	return &input{}
}

// Update updates the current input states.
func (i *input) Update() {
	x, y := ebiten.CursorPosition()
	// Set previous positions
	i.mousePrevPosX = i.mousePosX
	i.mousePrevPosY = i.mousePosY
	i.mousePrevGridPosX = i.mouseGridPosX
	i.mousePrevGridPosY = i.mouseGridPosY

	// Set current positions
	i.mousePosX = x
	i.mousePosY = y
	i.mouseGridPosX = x / pixelSize
	i.mouseGridPosY = y / pixelSize
	
	if ebiten.IsMouseButtonPressed(ebiten.MouseButtonRight) {
		i.mouseState = mouseStateRightClick
	} else if ebiten.IsMouseButtonPressed(ebiten.MouseButtonLeft) {
		i.mouseState = mouseStateLeftClick
	} else {
		i.mouseState = mouseStateNone
	}
	
	if ebiten.IsKeyPressed(ebiten.Key1) {
		i.mode = inputModeOutline
	} else if ebiten.IsKeyPressed(ebiten.Key2) {
		i.mode = inputModeSand
	} else if ebiten.IsKeyPressed(ebiten.Key3) {
		i.mode = inputModeWater
	} else if ebiten.IsKeyPressed(ebiten.Key4) {
		i.mode = inputModeWaterSpewer
	} else if ebiten.IsKeyPressed(ebiten.Key5) {
		i.mode = inputModeSandSpewer
	}
}