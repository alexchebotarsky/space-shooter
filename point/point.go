package point

import (
	"math"

	"github.com/goodleby/space-shooter/utils"
)

type Point struct {
	x float64
	y float64
}

func New(x, y float64) *Point {
	var p Point

	p.x = x
	p.y = y

	return &p
}

func (p *Point) Coordinates() (x, y float64) {
	return p.x, p.y
}

// MoveInDirection moves the point by the provided amount based on the
// direction. Direction is in range [0;1) with 0 being north, 0.25 - east, 0.5 -
// south and 0.75 - west. Note: direction can also be negative, in that case
// meaning of the values will be inverted as expected. This method is using
// triangular wave function to determine how much coordinates should be moved
// based on the direction:
// x(r)=4|(r-0.25)%1-0.5|-1
// y(r)=4|(r-0.5)%1-0.5|-1
func (p *Point) MoveInDirection(direction, amount float64) {
	p.x += (4*math.Abs(utils.EuclideanMod(direction-0.25, 1)-0.5) - 1) * amount
	p.y += (4*math.Abs(utils.EuclideanMod(direction-0.5, 1)-0.5) - 1) * amount
}
