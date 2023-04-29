package main

import (
	"fmt"
	"log"

	goebitencamerademo "github.com/co0p/ebiten-camera-demo"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
)

const Width, Height int = 800, 600

type Game04 struct {
	player     goebitencamerademo.Player
	background goebitencamerademo.Background
	world      goebitencamerademo.World

	mousePos goebitencamerademo.Position
}

func NewGame04() Game04 {
	return Game04{
		player:     goebitencamerademo.NewPlayer(100, 100),
		background: goebitencamerademo.NewBackground(0, 0),
		world:      goebitencamerademo.NewWorld(Width, Height),
	}
}

func (g *Game04) Update() error {

	x, y := ebiten.CursorPosition()
	g.mousePos.X = float64(x)
	g.mousePos.Y = float64(y)

	return g.player.Update()
}

func (g *Game04) Draw(screen *ebiten.Image) {

	g.background.Draw(g.world.Surface)
	g.player.Draw(g.world.Surface)

	worldDrawOptions := &ebiten.DrawImageOptions{}
	worldDrawOptions.GeoM.Translate(g.mousePos.X, g.mousePos.Y)
	screen.DrawImage(g.world.Surface, worldDrawOptions)

	g.printDebugInfo(screen)
}

func (g *Game04) printDebugInfo(screen *ebiten.Image) {
	screenBounds := screen.Bounds()
	debugStr := fmt.Sprintf("screen: %dx%d\nplayer: %s\nmouse: %s",
		screenBounds.Max.X, screenBounds.Max.Y,
		g.player.Position,
		g.mousePos,
	)
	ebitenutil.DebugPrint(screen, debugStr)
}

func (g *Game04) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return outsideWidth, outsideHeight
}

func main() {

	game := NewGame04()

	ebiten.SetFullscreen(false)
	ebiten.SetWindowSize(Width, Height)
	ebiten.SetWindowTitle("04 mouse")
	if err := ebiten.RunGame(&game); err != nil {
		log.Fatal(err)
	}
}
