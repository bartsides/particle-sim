package drawer

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
)

// Input represents the current key states.
type Input struct {
	mouseState		mouseState
	mousePosX		int
	mousePosY		int
	mousePrevPosX 	int
	mousePrevPosY 	int
	mode			inputMode
}

// NewInput generates a new Input object.
func NewInput() *Input {
	return &Input{}
}

// Update updates the current input states.
func (i *Input) Update() {
	x, y := ebiten.CursorPosition()
	i.mousePrevPosX = i.mousePosX
	i.mousePrevPosY = i.mousePosY
	i.mousePosX = x
	i.mousePosY = y
	
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
	}
}