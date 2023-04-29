package goebitencamerademo

import (
	"github.com/co0p/ebiten-camera-demo/assets"
	"github.com/hajimehoshi/ebiten/v2"
)

type Background struct {
	Position Position
	sprite   *ebiten.Image
}

func NewBackground(x, y float64) Background {
	return Background{
		Position: Position{X: x, Y: y},
		sprite:   assets.LoadBackgroundSprite(),
	}
}

func (b *Background) Draw(screen *ebiten.Image) {
	ops := ebiten.DrawImageOptions{}
	ops.GeoM.Scale(1.5, 1.5)
	ops.GeoM.Translate(b.Position.X, b.Position.Y)
	screen.DrawImage(b.sprite, &ops)
}

func (b *Background) Update() error { return nil }
