package hitbox

import (
	"image/color"

	"github.com/goodleby/space-shooter/point"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/vector"
)

type Hitbox struct {
	points   []*point.Point
	width    float64
	height   float64
	position *point.Point
}

func New(img *ebiten.Image, position *point.Point, points []*point.Point) *Hitbox {
	var p Hitbox

	p.points = points
	p.width = float64(img.Bounds().Dx())
	p.height = float64(img.Bounds().Dy())
	p.position = position

	return &p
}

func (h *Hitbox) Size() (float64, float64) {
	return h.width, h.height
}

func (h *Hitbox) Position() *point.Point {
	return h.position
}

func (h *Hitbox) CheckCollision(other *Hitbox) bool {
	return false
}

func (h *Hitbox) Draw(img *ebiten.Image) {
	imgWidth := float64(img.Bounds().Dx())
	imgHeight := float64(img.Bounds().Dy())

	x0, y0 := h.points[len(h.points)-1].Coordinates()
	for _, point := range h.points {
		x1, y1 := point.Coordinates()
		vector.StrokeLine(img, float32(x0*imgWidth), float32(y0*imgHeight), float32(x1*imgWidth), float32(y1*imgHeight), 1, color.RGBA{255, 0, 0, 255}, true)

		x0 = x1
		y0 = y1
	}
}
