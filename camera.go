package goebitencamerademo

import (
	"github.com/hajimehoshi/ebiten/v2"
)

const (
	zoomDefault = 1
	zoomFactor  = 0.02
	zoomMax     = 10
	zoomMin     = 0.01
)

type Camera struct {
	position Position
	zoom     float64
}

func NewCamera() Camera {
	return Camera{
		zoom: zoomDefault,
	}
}

func (c *Camera) SetPosition(p Position) {
	c.position = p
}

func (c *Camera) Move(p Position) {
	delta := p.Subtract(c.position)
	c.position = c.position.Add(delta.Scale(0.1))
}

func (c *Camera) ZoomIn() {
	z := c.zoom + zoomFactor
	c.SetZoom(z)
}

func (c *Camera) ZoomOut() {
	z := c.zoom - zoomFactor
	c.SetZoom(z)
}

func (c *Camera) SetZoom(z float64) {

	c.zoom = z

	if z < zoomMin {
		c.zoom = zoomMin
	}

	if z > zoomMax {
		c.zoom = zoomMax
	}
}

/*
	func (c *Camera) GetWorldPosition(screenPos Position) Position {
		return c.position.Add(screenPos)
		w, h := .Bounds().Dx(), screen.Bounds().Dy()
		co := math.Cos(-c.Rot)
		si := math.Sin(-c.Rot)

		x, y = (x-float64(w)/2)/c.Scale, (y-float64(h)/2)/c.Scale
		x, y = co*x-si*y, si*x+co*y

		return Position{X: x c.X, Y: y + c.Y}
	}
*/
func (c *Camera) Apply(world *ebiten.Image, screen *ebiten.Image) {
	ops := ebiten.DrawImageOptions{}
	w, h := screen.Bounds().Dx(), screen.Bounds().Dy()
	cx := float64(w) / 2.0
	cy := float64(h) / 2.0

	ops.GeoM.Translate(-cx, -cy)
	ops.GeoM.Scale(c.zoom, c.zoom)
	ops.GeoM.Translate(cx, cy)

	ops.GeoM.Translate(c.position.X*c.zoom, c.position.Y*c.zoom)
	screen.DrawImage(world, &ops)
}
