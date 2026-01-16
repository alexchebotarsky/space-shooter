package player

import (
	"github.com/goodleby/space-shooter/assets"
	"github.com/goodleby/space-shooter/hitbox"
	"github.com/goodleby/space-shooter/object"
	"github.com/goodleby/space-shooter/point"
	"github.com/goodleby/space-shooter/utils"
	"github.com/hajimehoshi/ebiten/v2"
)

type Player struct {
	assets *assets.Player

	position *point.Point
	rotation float64

	movementSpeed float64 // px per second
	rotationSpeed float64 // rotations per second

	hitbox *hitbox.Hitbox

	*object.Object
}

func New(assets *assets.Player) *Player {
	var p Player

	p.assets = assets

	windowWidth, windowHeight := ebiten.WindowSize()

	imgWidth := float64(assets.Ship.Bounds().Dx())
	imgHeight := float64(assets.Ship.Bounds().Dy())

	x := float64(windowWidth/2) - imgWidth/2
	y := float64(windowHeight/2) - imgHeight/2
	p.position = point.New(x, y)
	p.rotation = 0

	p.rotationSpeed = 0.75
	p.movementSpeed = 300

	p.hitbox = hitbox.New(assets.Ship, p.position, []*point.Point{
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
	})

	return &p
}

func (p *Player) Update() {
	rotationSpeed := p.rotationSpeed / float64(ebiten.TPS())
	movementSpeed := p.movementSpeed / float64(ebiten.TPS())

	if ebiten.IsKeyPressed(ebiten.KeyLeft) {
		p.rotation -= rotationSpeed
	}
	if ebiten.IsKeyPressed(ebiten.KeyRight) {
		p.rotation += rotationSpeed
	}
	p.rotation = utils.EuclideanMod(p.rotation, 1) // Keep rotation value within [0; 1) bounds

	if ebiten.IsKeyPressed(ebiten.KeyUp) {
		p.position.MoveInDirection(p.rotation, movementSpeed)
	}
}

func (p *Player) Draw(screen *ebiten.Image) {
	x, y := p.position.Coordinates()
	utils.DrawImage(screen, p.assets.Ship, x, y, 0.5, 0.5, p.rotation)
}

func (p *Player) Hitbox() *hitbox.Hitbox {
	return p.hitbox
}
