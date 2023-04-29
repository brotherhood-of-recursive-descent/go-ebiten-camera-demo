package goebitencamerademo

import "github.com/hajimehoshi/ebiten/v2"

type World struct {
	Width, Height int
	Surface       *ebiten.Image
}

func NewWorld(w, h int) World {
	return World{
		Width:   w,
		Height:  h,
		Surface: ebiten.NewImage(w, h),
	}
}
