package main

import (
	"log"

	"github.com/co0p/ebiten-camera-demo/camera"
	"github.com/co0p/ebiten-camera-demo/game"
	"github.com/hajimehoshi/ebiten/v2"
)

func main() {

	camera := camera.New()
	game := game.New(&camera)

	ebiten.SetFullscreen(true)
	ebiten.SetWindowTitle("Camera Demo")
	if err := ebiten.RunGame(&game); err != nil {
		log.Fatal(err)
	}
}
