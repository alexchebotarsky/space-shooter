package asteroid

import (
	"math/rand"

	"github.com/goodleby/space-shooter/point"
	"github.com/goodleby/space-shooter/utils"
	"github.com/hajimehoshi/ebiten/v2"
)

type Asteroid struct {
	img *ebiten.Image

	position *point.Point

	rotation      float64
	rotationSpeed float64

	movementDirection float64
	movementSpeed     float64
}

func New(img *ebiten.Image) *Asteroid {
	var a Asteroid

	a.img = img

	windowWidth, windowHeight := ebiten.WindowSize()

	imgWidth := float64(img.Bounds().Dx())
	imgHeight := float64(img.Bounds().Dy())

	// Choose random side
	switch rand.Intn(4) {
	case 0: // top
		a.position = point.New(utils.RandFloat(-imgWidth, float64(windowWidth)), -imgHeight)
		a.movementDirection = utils.RandFloat(0.25, 0.75) // west, south, east
	case 1: // bottom
		a.position = point.New(utils.RandFloat(-imgWidth, float64(windowWidth)), float64(windowHeight))
		a.movementDirection = utils.EuclideanMod(utils.RandFloat(0.75, 1.15), 1) // east, north, west. Note: 1.15 will become 0.15 after modulo
	case 2: // left
		a.position = point.New(-imgWidth, utils.RandFloat(-imgHeight, float64(windowHeight)))
		a.movementDirection = utils.RandFloat(0, 0.5) // north, west, south
	case 3: // right
		a.position = point.New(float64(windowWidth), utils.RandFloat(-imgHeight, float64(windowHeight)))
		a.movementDirection = utils.RandFloat(0.5, 1) // south, east, north
	}

	a.rotation = 0

	a.rotationSpeed = utils.RandFloat(-0.15, 0.15) // [-0.15; 0.15]
	a.movementSpeed = utils.RandFloat(50, 150)     // [50; 150]

	return &a
}

func (a *Asteroid) Update() {
	rotationSpeed := a.rotationSpeed / float64(ebiten.TPS())
	movementSpeed := a.movementSpeed / float64(ebiten.TPS())

	a.position.MoveInDirection(a.movementDirection, movementSpeed)

	a.rotation += rotationSpeed
	a.rotation = utils.EuclideanMod(a.rotation, 1) // Keep rotation value within [0; 1] bounds
}

func (a *Asteroid) Draw(screen *ebiten.Image) {
	x, y := a.position.Coordinates()
	utils.DrawImage(screen, a.img, x, y, 0.5, 0.5, a.rotation)
}

func (a *Asteroid) IsOutOfBounds() bool {
	windowWidth, windowHeight := ebiten.WindowSize()

	x, y := a.position.Coordinates()

	if x > float64(windowWidth)*1.5 || y > float64(windowHeight)*1.5 || x < -float64(windowWidth)*0.5 || y < -float64(windowHeight)*0.5 {
		return true
	}

	return false
}
