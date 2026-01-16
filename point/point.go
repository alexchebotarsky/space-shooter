package point

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

func (p *Point) X() float64 {
	return p.x
}

func (p *Point) Y() float64 {
	return p.y
}

func (p *Point) SetCoordinates(x, y float64) {
	p.x = x
	p.y = y
}

func (p *Point) SetX(x float64) {
	p.x = x
}

func (p *Point) SetY(y float64) {
	p.y = y
}
