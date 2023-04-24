package assets

import (
	"bytes"
	_ "embed"
	"image"
	_ "image/png"

	"github.com/hajimehoshi/ebiten/v2"
)

//go:embed background.png
var background []byte

//go:embed player.png
var player []byte

func LoadBackground() *ebiten.Image {
	src, _, _ := image.Decode(bytes.NewReader(background))
	return ebiten.NewImageFromImage(src)
}

func LoadPlayer() *ebiten.Image {
	src, _, _ := image.Decode(bytes.NewReader(player))
	return ebiten.NewImageFromImage(src)
}
