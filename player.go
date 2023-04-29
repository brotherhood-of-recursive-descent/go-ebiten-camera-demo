package goebitencamerademo

import (
	"github.com/co0p/ebiten-camera-demo/assets"
	"github.com/hajimehoshi/ebiten/v2"
)

const velocity float64 = 4

type Player struct {
	Position Position
	sprite   *ebiten.Image
}

func NewPlayer(x, y float64) Player {
	return Player{
		Position: Position{X: x, Y: y},
		sprite:   assets.LoadPlayerSprite(),
	}
}

func (p *Player) Draw(screen *ebiten.Image) {

	ops := ebiten.DrawImageOptions{}
	ops.GeoM.Translate(p.Position.X, p.Position.Y)
	screen.DrawImage(p.sprite, &ops)
}

func (p *Player) Update() error {
	if ebiten.IsKeyPressed(ebiten.KeyArrowLeft) {
		p.Position.X -= velocity
	}
	if ebiten.IsKeyPressed(ebiten.KeyArrowRight) {
		p.Position.X += velocity
	}
	if ebiten.IsKeyPressed(ebiten.KeyArrowUp) {
		p.Position.Y -= velocity
	}
	if ebiten.IsKeyPressed(ebiten.KeyArrowDown) {
		p.Position.Y += velocity
	}
	return nil
}
