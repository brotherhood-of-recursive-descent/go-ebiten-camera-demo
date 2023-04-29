package main

import (
	"fmt"
	"log"

	goebitencamerademo "github.com/co0p/ebiten-camera-demo"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
)

const Width, Height int = 800, 600

type Game03 struct {
	player     goebitencamerademo.Player
	background goebitencamerademo.Background
	world      goebitencamerademo.World
}

func NewGame03() Game03 {
	return Game03{
		player:     goebitencamerademo.NewPlayer(100, 100),
		background: goebitencamerademo.NewBackground(0, 0),
		world:      goebitencamerademo.NewWorld(Width, Height),
	}
}

func (g *Game03) Update() error {
	return g.player.Update()
}

func (g *Game03) Draw(screen *ebiten.Image) {

	g.background.Draw(g.world.Surface)
	g.player.Draw(g.world.Surface)

	worldDrawOptions := &ebiten.DrawImageOptions{}
	screen.DrawImage(g.world.Surface, worldDrawOptions)

	g.printDebugInfo(screen)
}

func (g *Game03) printDebugInfo(screen *ebiten.Image) {
	screenBounds := screen.Bounds()
	debugStr := fmt.Sprintf("screen: %dx%d\nplayer: %d,%d",
		screenBounds.Max.X, screenBounds.Max.Y,
		int(g.player.Position.X), int(g.player.Position.Y),
	)
	ebitenutil.DebugPrint(screen, debugStr)
}

func (g *Game03) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return outsideWidth, outsideHeight
}

func main() {

	game := NewGame03()

	ebiten.SetFullscreen(false)
	ebiten.SetWindowSize(Width, Height)
	ebiten.SetWindowTitle("03 world")
	if err := ebiten.RunGame(&game); err != nil {
		log.Fatal(err)
	}
}
