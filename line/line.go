package line

import "github.com/goodleby/space-shooter/point"

type Line struct {
	a *point.Point
	b *point.Point
}

func New(a, b *point.Point) *Line {
	var l Line

	l.a = a
	l.b = b

	return &l
}

func (v *Line) IsIntersecting(other *Line) bool {
	o1 := orientation(v, other.a)
	o2 := orientation(v, other.b)
	o3 := orientation(other, v.a)
	o4 := orientation(other, v.b)

	return o1*o2 < 0 && o3*o4 < 0
}

// orientation returns the orientation of the point p relative to the line l. If
// the point lies on the left side of the line, then a positive value is
// returned. If it lies on the right side, then a negative value is returned. If
// the point is collinear with the line, zero is returned.
func orientation(l *Line, p *point.Point) float64 {
	return (l.b.X()-l.a.X())*(p.Y()-l.a.Y()) - (l.b.Y()-l.a.Y())*(p.X()-l.a.X())
}
