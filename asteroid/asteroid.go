package asteroid

import (
	"math/rand"

	"github.com/goodleby/space-shooter/object"
	"github.com/goodleby/space-shooter/point"
	"github.com/goodleby/space-shooter/utils"
	"github.com/hajimehoshi/ebiten/v2"
)

type Asteroid struct {
	object            *object.Object
	movementDirection float64
	rotationSpeed     float64
	movementSpeed     float64
}

func New(img *ebiten.Image) *Asteroid {
	var a Asteroid

	windowWidth, windowHeight := ebiten.WindowSize()

	imgWidth := float64(img.Bounds().Dx())
	imgHeight := float64(img.Bounds().Dy())

	// Choose random side
	var x, y, movementDirection float64
	switch rand.Intn(4) {
	case 0: // top
		x = utils.RandFloat(-imgWidth, float64(windowWidth))
		y = -imgHeight
		movementDirection = utils.RandFloat(0.25, 0.75)
	case 1: // bottom
		x = utils.RandFloat(-imgWidth, float64(windowWidth))
		y = float64(windowHeight)
		movementDirection = utils.EuclideanMod(utils.RandFloat(0.75, 1.15), 1)
	case 2: // left
		x = -imgWidth
		y = utils.RandFloat(-imgHeight, float64(windowHeight))
		movementDirection = utils.RandFloat(0, 0.5)
	case 3: // right
		x = float64(windowWidth)
		y = utils.RandFloat(-imgHeight, float64(windowHeight))
		movementDirection = utils.RandFloat(0.5, 1)
	}

	hitpoints := []*point.Point{
		point.New(0.17, 0),
		point.New(0.73, 0),
		point.New(1, 0.49),
		point.New(0.85, 0.9),
		point.New(0.61, 0.84),
		point.New(0.29, 1),
		point.New(0, 0.61),
	}

	a.object = object.New(img, point.New(x, y), 0, hitpoints)
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

	topLeft := point.New(0, 0)
	bottomRight := point.New(float64(windowWidth), float64(windowHeight))

	return !a.object.IsWithinBounds(topLeft, bottomRight)
}
