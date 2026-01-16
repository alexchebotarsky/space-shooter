package vector

import "github.com/goodleby/space-shooter/point"

type Vector struct {
	from *point.Point
	to   *point.Point
}

func New(from, to *point.Point) *Vector {
	var v Vector

	v.from = from
	v.to = to

	return &v
}

func (v *Vector) IsIntersecting(other *Vector) bool {
	o1 := orientation(v, other.from)
	o2 := orientation(v, other.to)
	o3 := orientation(other, v.from)
	o4 := orientation(other, v.to)

	return o1*o2 < 0 && o3*o4 < 0
}

// orientation returns the orientation of the point p relative to the vector v.
// If the point lies on the left side of the vector, then a positive value is
// returned. If it lies on the right side, then a negative value is returned. If
// the point is collinear with the vector, zero is returned.
func orientation(v *Vector, p *point.Point) float64 {
	return (v.to.X()-v.from.X())*(p.Y()-v.from.Y()) - (v.to.Y()-v.from.Y())*(p.X()-v.from.X())
}
