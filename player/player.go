package player

import (
	"github.com/goodleby/space-shooter/assets"
	"github.com/goodleby/space-shooter/object"
	"github.com/goodleby/space-shooter/point"
	"github.com/hajimehoshi/ebiten/v2"
)

type Player struct {
	assets        *assets.Player
	object        *object.Object
	movementSpeed float64 // px per second
	rotationSpeed float64 // rotations per second
}

func New(assets *assets.Player) *Player {
	var p Player

	windowWidth, windowHeight := ebiten.WindowSize()

	imgWidth := float64(assets.Ship.Bounds().Dx())
	imgHeight := float64(assets.Ship.Bounds().Dy())

	x := float64(windowWidth/2) - imgWidth/2
	y := float64(windowHeight/2) - imgHeight/2

	hitpoints := []*point.Point{
		point.New(0.42, 0),
		point.New(0.58, 0),
		point.New(0.62, 0.34),
		point.New(0.87, 0.54),
		point.New(1, 0.37),
		point.New(0.95, 0.86),
		point.New(0.91, 0.78),
		point.New(0.63, 0.87),
		point.New(0.57, 1),
		point.New(0.43, 1),
		point.New(0.37, 0.87),
		point.New(0.09, 0.78),
		point.New(0.05, 0.86),
		point.New(0, 0.37),
		point.New(0.13, 0.54),
		point.New(0.38, 0.34),
	}

	p.assets = assets
	p.object = object.New(assets.Ship, point.New(x, y), 0, hitpoints)
	p.movementSpeed = 300
	p.rotationSpeed = 0.75

	return &p
}

func (p *Player) Update() {
	rotationSpeed := p.rotationSpeed / float64(ebiten.TPS())
	movementSpeed := p.movementSpeed / float64(ebiten.TPS())

	if ebiten.IsKeyPressed(ebiten.KeyLeft) {
		p.object.RotateBy(-rotationSpeed)
	}
	if ebiten.IsKeyPressed(ebiten.KeyRight) {
		p.object.RotateBy(rotationSpeed)
	}

	if ebiten.IsKeyPressed(ebiten.KeyUp) {
		p.object.MoveBy(movementSpeed)
	}
}

func (p *Player) Draw(screen *ebiten.Image) {
	p.object.Draw(screen)
}

func (p *Player) IsWithinBounds() bool {
	windowWidth, windowHeight := ebiten.WindowSize()

	topLeft := point.New(0, 0)
	bottomRight := point.New(float64(windowWidth), float64(windowHeight))

	return p.object.IsWithinBounds(topLeft, bottomRight)
}
