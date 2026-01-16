package object

import (
	"fmt"
	"image/color"
	"math"

	"github.com/goodleby/space-shooter/point"
	"github.com/goodleby/space-shooter/utils"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"github.com/hajimehoshi/ebiten/v2/vector"
)

type Object struct {
	img       *ebiten.Image
	position  *point.Point
	rotation  float64
	hitpoints []*point.Point
}

func New(image *ebiten.Image, position *point.Point, rotation float64, hitpoints []*point.Point) *Object {
	var o Object

	o.img = image
	o.position = position
	o.rotation = rotation
	o.hitpoints = hitpoints

	return &o
}

// func (o1 *Object) CheckCollision(o2 *Object) bool {
// 	// var vA, vB *vector.Vector

// 	p1 := o1.hitpoints[len(o1.hitpoints)-1]
// 	for _, p2 := range o1.hitpoints {
// 		// vA = vector.New(p1, p2)

// 		log.Printf("Vector A: (%.1f,%.1f)-(%.1f,%.1f)", p1.X(), p1.Y(), p2.X(), p2.Y())

// 		p1 = p2
// 	}

// 	return false
// }

func (o *Object) IsWithinBounds(a, b *point.Point) bool {
	posX, posY := o.position.Coordinates()
	imgWidth := float64(o.img.Bounds().Dx())
	imgHeight := float64(o.img.Bounds().Dy())

	for _, hp := range o.hitpoints {
		x, y := utils.GetOriginPoint(posX, posY, imgWidth, imgHeight, o.rotation, hp.X(), hp.Y())

		if x < a.X() || y < a.Y() || x > b.X() || y > b.Y() {
			return false
		}
	}

	return true
}

func (o *Object) Draw(screen *ebiten.Image) {
	posX, posY := o.position.Coordinates()
	windowWidth, windowHeight := ebiten.WindowSize()
	imgWidth := float64(o.img.Bounds().Dx())
	imgHeight := float64(o.img.Bounds().Dy())

	utils.DrawImage(screen, o.img, posX, posY, o.rotation, 0.5, 0.5)

	a := o.hitpoints[len(o.hitpoints)-1]
	for _, b := range o.hitpoints {
		ax, ay := utils.GetOriginPoint(posX, posY, imgWidth, imgHeight, o.rotation, a.X(), a.Y())
		bx, by := utils.GetOriginPoint(posX, posY, imgWidth, imgHeight, o.rotation, b.X(), b.Y())

		vector.StrokeLine(screen, float32(ax), float32(ay), float32(bx), float32(by), 1, color.RGBA{0, 255, 0, 255}, true)

		a = b
	}

	isWithinBounds := o.IsWithinBounds(point.New(0, 0), point.New(float64(windowWidth), float64(windowHeight)))
	ebitenutil.DebugPrint(screen, fmt.Sprintf("position(x=%.1f,y=%.1f), rotation(%.2f), isWithinBounds(%v)", posX, posY, o.rotation, isWithinBounds))
}

func (o *Object) MoveBy(amount float64) {
	o.MoveInDirection(o.rotation, amount)
}

// MoveInDirection moves the point by the provided amount based on the
// direction. Direction is in range [0;1) with 0 being north, 0.25 - east, 0.5 -
// south and 0.75 - west. Note: direction can also be negative, in that case
// meaning of the values will be inverted as expected. This method is using
// triangular wave function to determine how much coordinates should be moved
// based on the direction:
// x(r)=4|(r-0.25)%1-0.5|-1
// y(r)=4|(r-0.5)%1-0.5|-1
func (o *Object) MoveInDirection(direction, amount float64) {
	o.position.SetX(o.position.X() + (4*math.Abs(utils.EuclideanMod(direction-0.25, 1)-0.5)-1)*amount)
	o.position.SetY(o.position.Y() + (4*math.Abs(utils.EuclideanMod(direction-0.5, 1)-0.5)-1)*amount)
}

func (o *Object) RotateBy(amount float64) {
	// Update rotation and keep it within [0; 1) bounds
	o.rotation = utils.EuclideanMod(o.rotation+amount, 1)
}
