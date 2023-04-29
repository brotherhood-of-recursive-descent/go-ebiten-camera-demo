package main

import (
	"fmt"
	"log"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
)

type Game struct{}

func (g *Game) Update() error { return nil }

func (g *Game) Draw(screen *ebiten.Image) {
	screenBounds := screen.Bounds()
	debugStr := fmt.Sprintf("screen: %dx%d", screenBounds.Max.X, screenBounds.Max.Y)
	ebitenutil.DebugPrint(screen, debugStr)
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return outsideWidth, outsideHeight
}

func main() {

	game := &Game{}

	ebiten.SetFullscreen(false)
	ebiten.SetWindowTitle("01 screen")
	if err := ebiten.RunGame(game); err != nil {
		log.Fatal(err)
	}
}
