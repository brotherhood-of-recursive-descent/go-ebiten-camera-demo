package main

import (
	"fmt"
	"log"

	goebitencamerademo "github.com/co0p/ebiten-camera-demo"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
)

const ScreenWidth, ScreenHeight int = 800, 600

type Game07 struct {
	player     goebitencamerademo.Player
	background goebitencamerademo.Background
	world      goebitencamerademo.World
	camera     goebitencamerademo.Camera

	mousePos goebitencamerademo.Position
}

func NewGame07() Game07 {
	return Game07{
		world:      goebitencamerademo.NewWorld(ScreenWidth, ScreenHeight),
		player:     goebitencamerademo.NewPlayer(100, 100),
		background: goebitencamerademo.NewBackground(0, 0),
		camera:     goebitencamerademo.NewCamera(),
	}
}

func (g *Game07) Update() error {

	x, y := ebiten.CursorPosition()
	g.mousePos.X = float64(x)
	g.mousePos.Y = float64(y)

	return g.player.Update()
}

func (g *Game07) Draw(screen *ebiten.Image) {

	g.background.Draw(g.world.Surface)
	g.player.Draw(g.world.Surface)

	x := float64(ScreenWidth/2) - g.player.Position.X
	y := float64(ScreenHeight/2) - g.player.Position.Y

	g.camera.SetPosition(x, y)
	g.camera.Apply(g.world.Surface, screen)

	g.printDebugInfo(screen)
}

func (g *Game07) printDebugInfo(screen *ebiten.Image) {
	screenBounds := screen.Bounds()
	debugStr := fmt.Sprintf("screen: %dx%d\nplayer: %s\nmouse: %s",
		screenBounds.Max.X, screenBounds.Max.Y,
		g.player.Position,
		g.mousePos,
	)
	ebitenutil.DebugPrint(screen, debugStr)
}

func (g *Game07) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return outsideWidth, outsideHeight
}

func main() {

	game := NewGame07()

	ebiten.SetFullscreen(false)
	ebiten.SetWindowSize(ScreenWidth, ScreenHeight)
	ebiten.SetWindowTitle("06 center")
	if err := ebiten.RunGame(&game); err != nil {
		log.Fatal(err)
	}
}
