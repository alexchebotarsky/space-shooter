package object

import (
	"github.com/goodleby/space-shooter/point"
	"github.com/goodleby/space-shooter/utils"
	"github.com/goodleby/space-shooter/vector"
	"github.com/hajimehoshi/ebiten/v2"
)

type Object struct {
	img       *ebiten.Image
	hitpoints []*point.Point
	position  *point.Point
	rotation  float64
}

func New(image *ebiten.Image, hitpoints []*point.Point, position *point.Point, rotation float64) *Object {
	var o Object

	o.img = image
	o.hitpoints = hitpoints
	o.position = position
	o.rotation = rotation

	return &o
}

func (o *Object) CheckCollision(other *Object) bool {
	var v1, v2 *vector.Vector
	for _, p := range p.hitpoints {
		v1 = vector.New(o.hitpoints[len(o.hitpoints)-1])
	}
}

func (o *Object) Draw(screen *ebiten.Image) {
	x, y := o.position.Coordinates()
	utils.DrawImage(screen, o.img, x, y, 0.5, 0.5, o.rotation)
}
