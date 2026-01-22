package player

import (
	"github.com/goodleby/space-shooter/assets"
	"github.com/goodleby/space-shooter/object"
	"github.com/goodleby/space-shooter/point"
	"github.com/goodleby/space-shooter/utils"
	"github.com/hajimehoshi/ebiten/v2"
)

type Player struct {
	assets        *assets.Player
	object        *object.Object
	movementSpeed float64 // px per second
	rotationSpeed float64 // rotations per second
}

func New(x, y float64, assets *assets.Player) *Player {
	var p Player

	imgWidth, imgHeight := utils.ImageSize(assets.Ship)
	imgHitpoints := []*point.Point{
		point.New(imgWidth*0.42, imgHeight*0.00),
		point.New(imgWidth*0.58, imgHeight*0.00),
		point.New(imgWidth*0.62, imgHeight*0.34),
		point.New(imgWidth*0.87, imgHeight*0.54),
		point.New(imgWidth*1.00, imgHeight*0.37),
		point.New(imgWidth*0.95, imgHeight*0.86),
		point.New(imgWidth*0.91, imgHeight*0.78),
		point.New(imgWidth*0.63, imgHeight*0.87),
		point.New(imgWidth*0.57, imgHeight*1.00),
		point.New(imgWidth*0.43, imgHeight*1.00),
		point.New(imgWidth*0.37, imgHeight*0.87),
		point.New(imgWidth*0.09, imgHeight*0.78),
		point.New(imgWidth*0.05, imgHeight*0.86),
		point.New(imgWidth*0.00, imgHeight*0.37),
		point.New(imgWidth*0.13, imgHeight*0.54),
		point.New(imgWidth*0.38, imgHeight*0.34),
	}

	p.assets = assets
	p.object = object.New(point.New(x, y), 0, assets.Ship, imgHitpoints)
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

func (p *Player) IsIntersecting(object *object.Object) bool {
	return p.object.IsIntersecting(object)
}
