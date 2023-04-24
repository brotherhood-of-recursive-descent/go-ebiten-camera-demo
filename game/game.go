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
}

func (g *Game) Update() error {

	g.handlePlayerInput()

	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	// draw the background
	opsB := ebiten.DrawImageOptions{}
	opsB.GeoM.Scale(3, 3)
	screen.DrawImage(g.background, &opsB)

	// draw the player
	opsP := ebiten.DrawImageOptions{}
	opsP.GeoM.Scale(g.scale, g.scale)
	opsP.GeoM.Translate(g.posX, g.posY)
	screen.DrawImage(g.player, &opsP)
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return outsideWidth, outsideHeight
}
