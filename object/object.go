package object

import (
	"image/color"

	"github.com/goodleby/space-shooter/line"
	"github.com/goodleby/space-shooter/point"
	"github.com/goodleby/space-shooter/utils"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/vector"
)

type Object struct {
	position     *point.Point
	rotation     float64
	img          *ebiten.Image
	imgHitpoints []*point.Point
}

// New creates new object instance
func New(position *point.Point, rotation float64, img *ebiten.Image, imgHitpoints []*point.Point) *Object {
	var o Object

	o.position = position
	o.rotation = rotation
	o.img = img
	o.imgHitpoints = imgHitpoints

	return &o
}

func (o *Object) Hitpoints() []*point.Point {
	hitpoints := make([]*point.Point, 0, len(o.imgHitpoints))

	posX, posY := o.position.Coordinates()
	width, height := utils.ImageSize(o.img)

	for _, p := range o.imgHitpoints {
		localX, localY := utils.TransformedLocalPoint(p.X(), p.Y(), width, height, o.rotation, 0.5, 0.5)
		// Absolute position + Local position - Half of the size (to center)
		hitpoints = append(hitpoints, point.New(posX+localX-width/2, posY+localY-height/2))
	}

	return hitpoints
}

func (o *Object) Hitlines() []*line.Line {
	points := o.Hitpoints()
	lines := make([]*line.Line, 0, len(points)+1)

	a := points[len(points)-1]
	for _, b := range points {
		lines = append(lines, line.New(a, b))
		a = b
	}

	return lines
}

// IsIntersecting checks whether one object intersects with the other.
func (o1 *Object) IsIntersecting(o2 *Object) bool {
	hitlines1 := o1.Hitlines()
	hitlines2 := o2.Hitlines()

	for _, ab := range hitlines1 {
		for _, cd := range hitlines2 {
			if ab.IsIntersecting(cd) {
				return true
			}
		}
	}

	return false
}

func (o *Object) IsOutOfBounds(min, max *point.Point) bool {
	for _, hp := range o.Hitpoints() {
		if !(hp.X() < min.X() || hp.X() > max.X() || hp.Y() < min.Y() || hp.Y() > max.Y()) {
			return false
		}
	}

	return true
}

// Draw draws object's image on the screen with all transformations applied.
func (o *Object) Draw(screen *ebiten.Image) {
	x, y := o.position.Coordinates()
	utils.DrawImage(screen, o.img, x, y, o.rotation, 0.5, 0.5)

	o.DrawHitlines(screen)
}

func (o *Object) DrawHitlines(screen *ebiten.Image) {
	hitpoints := o.Hitpoints()

	a := hitpoints[len(hitpoints)-1]
	for _, b := range hitpoints {
		vector.StrokeLine(screen, float32(a.X()), float32(a.Y()), float32(b.X()), float32(b.Y()), 2, color.RGBA{0, 255, 0, 0}, true)
		a = b
	}
}

// MoveInDirection moves the object by an amount in the given direction.
func (o *Object) MoveInDirection(direction, amount float64) {
	mx, my := utils.DirectionMultipliers(direction)
	o.position.SetX(o.position.X() + amount*mx)
	o.position.SetY(o.position.Y() + amount*my)
}

// MoveBy moves the object by an amount in the direction it's pointing at.
func (o *Object) MoveBy(amount float64) {
	o.MoveInDirection(o.rotation, amount)
}

// RotateBy rotates counter clockwise the object by the given amount.
func (o *Object) RotateBy(amount float64) {
	// Mod here is to keep the value within [0; 1) bounds
	o.rotation = utils.EuclideanMod(o.rotation+amount, 1)
}
