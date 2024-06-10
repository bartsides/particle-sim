package main

import (
	"log"

	"github.com/bartsides/particle-sim/particle"
	"github.com/hajimehoshi/ebiten/v2"
)

/*
TODO: Dynamically change canvas size on window resize?

*/

func main() {
	canvas, err := particle.New()
	if err != nil {
		log.Fatal(err)
	}
	ebiten.SetWindowSize(particle.Width, particle.Height)
	ebiten.SetWindowTitle("Particle Sim")
	if err := ebiten.RunGame(canvas); err != nil {
		log.Fatal(err)
	}
}
