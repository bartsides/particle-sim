package drawer

import (
	"image/color"
	"math/rand"
)

var (
	backgroundColor = color.RGBA{250, 248, 239, 255}
	outlineColor	= color.RGBA{187, 173, 160, 255}
	sandColors		= []color.RGBA{
		{237, 188, 121, 255},
		{230, 198, 101, 255},
	}
	waterColors		= []color.RGBA{
		{ 30, 129, 176, 255},
		{ 30, 100, 176, 255},
		{ 60, 120, 176, 255},
		{ 40, 120, 215, 255},
	}
)

func getRandomColor(colors []color.RGBA) color.RGBA {
	return colors[rand.Intn(len(colors))]
}