package player

import (
	"time"

	"github.com/goodleby/space-shooter/assets"
	"github.com/goodleby/space-shooter/bullet"
	"github.com/goodleby/space-shooter/object"
	"github.com/goodleby/space-shooter/point"
	"github.com/goodleby/space-shooter/timer"
	"github.com/goodleby/space-shooter/utils"
	"github.com/hajimehoshi/ebiten/v2"
)

type Player struct {
	assets        *assets.Player
	object        *object.Object
	movementSpeed float64 // px per second
	rotationSpeed float64 // rotations per second

	bulletCooldown *timer.Timer
	bullets        []*bullet.Bullet
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
	p.rotationSpeed = 0.5

	p.bulletCooldown = timer.New(500 * time.Millisecond)
	p.bullets = []*bullet.Bullet{}

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

	if ebiten.IsKeyPressed(ebiten.KeySpace) {
		if p.bulletCooldown.IsReady() {
			x, y := p.object.Coordinates()
			shipWidth, shipHeight := utils.ImageSize(p.assets.Ship)
			bx, by := utils.TransformedLocalPoint(shipWidth/2, 0, shipWidth, shipHeight, p.object.Rotation(), 0.5, 0.5)
			bullet := bullet.New(point.New(x+bx-shipWidth/2, y+by-shipHeight/2), p.object.Rotation(), p.assets.Laser)
			p.bullets = append(p.bullets, bullet)
			p.bulletCooldown.Reset()
		}
	}

	p.bulletCooldown.Update()

	for i, bullet := range p.bullets {
		if bullet.IsOutOfBounds() {
			p.bullets = append(p.bullets[:i], p.bullets[i+1:]...)
			continue
		}

		bullet.Update()
	}
}

func (p *Player) Draw(screen *ebiten.Image) {
	for _, bullet := range p.bullets {
		bullet.Draw(screen)
	}

	p.object.Draw(screen)
}

func (p *Player) IsIntersecting(object *object.Object) bool {
	return p.object.IsIntersecting(object)
}

func (p *Player) HasHit(object *object.Object) bool {
	for i, bullet := range p.bullets {
		if bullet.IsIntersecting(object) {
			p.bullets = append(p.bullets[:i], p.bullets[i+1:]...)
			return true
		}
	}

	return false
}
