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

func (p Position) Add(o Position) Position {
	return Position{
		p.X + o.X,
		p.Y + o.Y,
	}
}

func (p Position) Subtract(o Position) Position {
	return Position{
		p.X - o.X,
		p.Y - o.Y,
	}
}

func (p Position) Scale(s float64) Position {
	return Position{
		p.X * s,
		p.Y * s,
	}
}
func (p Position) String() string {
	return fmt.Sprintf("%d,%d", int(p.X), int(p.Y))
}
