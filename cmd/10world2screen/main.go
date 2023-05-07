package main

import (
	"fmt"
	"log"

	goebitencamerademo "github.com/co0p/ebiten-camera-demo"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
)

const ScreenWidth, ScreenHeight int = 800, 600

type Game10 struct {
	player     goebitencamerademo.Player
	background goebitencamerademo.Background
	world      goebitencamerademo.World
	camera     goebitencamerademo.Camera

	mousePos goebitencamerademo.Position
}

func NewGame10() Game10 {
	return Game10{
		world:      goebitencamerademo.NewWorld(ScreenWidth*2, ScreenHeight*2),
		player:     goebitencamerademo.NewPlayer(100, 100),
		background: goebitencamerademo.NewBackground(0, 0),
		camera:     goebitencamerademo.NewCamera(),
	}
}

func (g *Game10) Update() error {

	x, y := ebiten.CursorPosition()
	g.mousePos.X = float64(x)
	g.mousePos.Y = float64(y)

	if ebiten.IsKeyPressed(ebiten.KeyPageUp) {
		g.camera.ZoomIn()
	}
	if ebiten.IsKeyPressed(ebiten.KeyPageDown) {
		g.camera.ZoomOut()
	}

	if ebiten.IsMouseButtonPressed(ebiten.MouseButton0) {
		mouseInWorld := g.camera.GetWorldPosition(g.mousePos)
		g.camera.SetPosition(mouseInWorld)
	}

	return g.player.Update()
}

func (g *Game10) Draw(screen *ebiten.Image) {

	g.background.Draw(g.world.Surface)
	g.player.Draw(g.world.Surface)

	g.camera.Apply(g.world.Surface, screen)

	g.printDebugInfo(screen)
}

func (g *Game10) printDebugInfo(screen *ebiten.Image) {
	screenBounds := screen.Bounds()
	debugStr := fmt.Sprintf("screen: %dx%d\nplayer: %s\nmouse: %s",
		screenBounds.Max.X, screenBounds.Max.Y,
		g.player.Position,
		g.mousePos,
	)
	ebitenutil.DebugPrint(screen, debugStr)
}

func (g *Game10) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return outsideWidth, outsideHeight
}

func main() {

	game := NewGame10()

	ebiten.SetFullscreen(false)
	ebiten.SetWindowSize(ScreenWidth, ScreenHeight)
	ebiten.SetWindowTitle("06 center")
	if err := ebiten.RunGame(&game); err != nil {
		log.Fatal(err)
	}
}
