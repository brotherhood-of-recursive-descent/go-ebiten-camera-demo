package game

import (
	"github.com/co0p/ebiten-camera-demo/assets"
	"github.com/co0p/ebiten-camera-demo/camera"
	"github.com/hajimehoshi/ebiten/v2"
)

type Game struct {
	player     *ebiten.Image
	background *ebiten.Image

	posX, posY, scale float64
	playerSpeed       float64

	mouseX, mouseY float64

	world *ebiten.Image
}

func New(camera *camera.Camera) Game {

	return Game{

		background: assets.LoadBackground(),
		player:     assets.LoadPlayer(),

		scale:       1.5,
		posX:        400,
		posY:        500,
		playerSpeed: 5,
	}

}

func (g *Game) handlePlayerInput() {
	// player input
	if ebiten.IsKeyPressed(ebiten.KeyArrowLeft) {
		g.posX -= g.playerSpeed
	}
	if ebiten.IsKeyPressed(ebiten.KeyArrowRight) {
		g.posX += g.playerSpeed
	}
	if ebiten.IsKeyPressed(ebiten.KeyArrowUp) {
		g.posY -= g.playerSpeed
	}
	if ebiten.IsKeyPressed(ebiten.KeyArrowDown) {
		g.posY += g.playerSpeed
	}

	// save mouse position
	x, y := ebiten.CursorPosition()
	g.mouseX, g.mouseY = float64(x), float64(y)
}

func (g *Game) Update() error {

	g.handlePlayerInput()

	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {

	screenBounds := screen.Bounds()

	// lazy initializing the world to render to, as image
	// creation is very expensive
	if g.world == nil {
		g.world = ebiten.NewImage(screenBounds.Dx(), screenBounds.Dy())
	}

	// draw the background
	opsB := ebiten.DrawImageOptions{}
	opsB.GeoM.Scale(3, 3)
	g.world.DrawImage(g.background, &opsB)

	// draw the player
	opsP := ebiten.DrawImageOptions{}
	opsP.GeoM.Scale(g.scale, g.scale)
	opsP.GeoM.Translate(g.posX, g.posY)
	g.world.DrawImage(g.player, &opsP)

	// draw the world onto screen
	worldDrawOptions := &ebiten.DrawImageOptions{}
	worldDrawOptions.GeoM.Translate(g.mouseX-float64(screenBounds.Dx())/2, g.mouseY-float64(screenBounds.Dy())/2)
	screen.DrawImage(g.world, worldDrawOptions)
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return outsideWidth, outsideHeight
}
