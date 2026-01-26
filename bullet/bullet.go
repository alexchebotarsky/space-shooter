package bullet

import (
	"github.com/goodleby/space-shooter/object"
	"github.com/goodleby/space-shooter/point"
	"github.com/goodleby/space-shooter/utils"
	"github.com/hajimehoshi/ebiten/v2"
)

type Bullet struct {
	object        *object.Object
	img           *ebiten.Image
	movementSpeed float64
}

func New(position *point.Point, direction float64, img *ebiten.Image) *Bullet {
	var b Bullet

	imgWidth, imgHeight := utils.ImageSize(img)
	imgHitpoints := []*point.Point{
		point.New(imgWidth*0.15, imgHeight*0.05),
		point.New(imgWidth*0.85, imgHeight*0.05),
		point.New(imgWidth*0.85, imgHeight*0.55),
		point.New(imgWidth*0.15, imgHeight*0.55),
	}

	b.object = object.New(position, direction, img, imgHitpoints)
	b.img = img
	b.movementSpeed = 1000

	return &b
}

func (b *Bullet) Update() {
	movementSpeed := b.movementSpeed / float64(ebiten.TPS())
	b.object.MoveBy(movementSpeed)
}

func (b *Bullet) Draw(screen *ebiten.Image) {
	b.object.Draw(screen)
}

func (b *Bullet) IsOutOfBounds() bool {
	windowWidth, windowHeight := ebiten.WindowSize()

	return b.object.IsOutOfBounds(
		point.New(0, 0),
		point.New(float64(windowWidth), float64(windowHeight)),
	)
}

func (b *Bullet) IsIntersecting(object *object.Object) bool {
	return b.object.IsIntersecting(object)
}
