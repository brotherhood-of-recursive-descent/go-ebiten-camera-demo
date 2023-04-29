package main

import (
	"fmt"
	"log"

	goebitencamerademo "github.com/co0p/ebiten-camera-demo"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
)

type Game02 struct {
	player     goebitencamerademo.Player
	background goebitencamerademo.Background
}

func NewGame02() Game02 {
	return Game02{
		player:     goebitencamerademo.NewPlayer(100, 100),
		background: goebitencamerademo.NewBackground(0, 0),
	}
}

func (g *Game02) Update() error {
	return g.player.Update()
}

func (g *Game02) Draw(screen *ebiten.Image) {

	g.background.Draw(screen)
	g.player.Draw(screen)

	g.printDebugInfo(screen)
}

func (g *Game02) printDebugInfo(screen *ebiten.Image) {
	screenBounds := screen.Bounds()
	debugStr := fmt.Sprintf("screen: %dx%d\nplayer: %d,%d",
		screenBounds.Max.X, screenBounds.Max.Y,
		int(g.player.Position.X), int(g.player.Position.Y),
	)
	ebitenutil.DebugPrint(screen, debugStr)
}

func (g *Game02) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return outsideWidth, outsideHeight
}

func main() {

	game := NewGame02()

	ebiten.SetFullscreen(false)
	ebiten.SetWindowTitle("02 objects")
	if err := ebiten.RunGame(&game); err != nil {
		log.Fatal(err)
	}
}
