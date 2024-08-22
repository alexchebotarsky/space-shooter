package player

import (
	"github.com/goodleby/space-shooter/assets"
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
