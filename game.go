package goebitencamerademo

import (
	"fmt"

	"github.com/hajimehoshi/ebiten/v2"
)

type GameObject interface {
	Draw(screen *ebiten.Image)
	Update() error
}

type Position struct {
	X, Y float64
}

func (p Position) String() string {
	return fmt.Sprintf("%d,%d", int(p.X), int(p.Y))
}
