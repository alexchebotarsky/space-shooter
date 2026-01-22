package asteroid

import (
	"math/rand"

	"github.com/goodleby/space-shooter/object"
	"github.com/goodleby/space-shooter/point"
	"github.com/goodleby/space-shooter/utils"
	"github.com/hajimehoshi/ebiten/v2"
)

type Asteroid struct {
	img               *ebiten.Image
	object            *object.Object
	movementDirection float64
	rotationSpeed     float64
	movementSpeed     float64
}

func New(img *ebiten.Image) *Asteroid {
	var a Asteroid

	windowWidth, windowHeight := ebiten.WindowSize()
	imgWidth, imgHeight := utils.ImageSize(img)

	// Choose random side
	var x, y, movementDirection float64
	switch rand.Intn(4) {
	case 0: // top
		x = utils.RandFloat(-imgWidth, float64(windowWidth)+imgWidth)
		y = -imgHeight
		movementDirection = utils.RandFloat(0.25, 0.75)
	case 1: // bottom
		x = utils.RandFloat(-imgWidth, float64(windowWidth)+imgWidth)
		y = float64(windowHeight) + imgHeight
		movementDirection = utils.EuclideanMod(utils.RandFloat(0.75, 1.15), 1)
	case 2: // left
		x = -imgWidth
		y = utils.RandFloat(-imgHeight, float64(windowHeight)+imgHeight)
		movementDirection = utils.RandFloat(0, 0.5)
	case 3: // right
		x = float64(windowWidth) + imgWidth
		y = utils.RandFloat(-imgHeight, float64(windowHeight)+imgHeight)
		movementDirection = utils.RandFloat(0.5, 1)
	}

	imgHitpoints := []*point.Point{
		point.New(imgWidth*0.17, imgHeight*0.00),
		point.New(imgWidth*0.73, imgHeight*0.00),
		point.New(imgWidth*1.00, imgHeight*0.49),
		point.New(imgWidth*0.85, imgHeight*0.90),
		point.New(imgWidth*0.61, imgHeight*0.84),
		point.New(imgWidth*0.29, imgHeight*1.00),
		point.New(imgWidth*0.00, imgHeight*0.61),
	}

	a.img = img
	a.object = object.New(point.New(x, y), 0, img, imgHitpoints)
	a.movementDirection = movementDirection
	a.rotationSpeed = utils.RandFloat(-0.15, 0.15) // [-0.15; 0.15]
	a.movementSpeed = utils.RandFloat(50, 150)     // [50; 150]

	return &a
}

func (a *Asteroid) Update() {
	rotationSpeed := a.rotationSpeed / float64(ebiten.TPS())
	movementSpeed := a.movementSpeed / float64(ebiten.TPS())

	a.object.MoveInDirection(a.movementDirection, movementSpeed)
	a.object.RotateBy(rotationSpeed)
}

func (a *Asteroid) Draw(screen *ebiten.Image) {
	a.object.Draw(screen)
}

func (a *Asteroid) IsOutOfBounds() bool {
	windowWidth, windowHeight := ebiten.WindowSize()
	imgWidth, imgHeight := utils.ImageSize(a.img)

	return a.object.IsOutOfBounds(
		point.New(-imgWidth, -imgHeight),
		point.New(float64(windowWidth)+imgWidth, float64(windowHeight)+imgHeight),
	)
}

func (a *Asteroid) Object() *object.Object {
	return a.object
}
