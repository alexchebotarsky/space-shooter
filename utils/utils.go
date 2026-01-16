package utils

import (
	"math"
	"math/rand"

	"github.com/hajimehoshi/ebiten/v2"
)

func RandFloat(min, max float64) float64 {
	return rand.Float64()*(max-min) + min
}

// EuclideanMod is similar to math.Mod, but the result is always positive. Read
// about different types of mod: https://en.wikipedia.org/wiki/Modulo
func EuclideanMod(x, y float64) float64 {
	y = math.Abs(y)
	return math.Mod(math.Mod(x, y)+y, y)
}

// FlooredMod is similar to math.Mod, but the result sign is that of y instead
// of x. Read about different types of mod: https://en.wikipedia.org/wiki/Modulo
func FlooredMod(x, y float64) float64 {
	return math.Mod(math.Mod(x, y)+y, y)
}

// DrawImage is similar to ebiten.DrawImage, but allows specifying origin point
// in [0;1] range rotation in full rotations.
func DrawImage(screen *ebiten.Image, img *ebiten.Image, x, y, rotation, originX, originY float64) {
	options := &ebiten.DrawImageOptions{}

	w := float64(img.Bounds().Dx())
	h := float64(img.Bounds().Dy())

	tx := x - (w * originX)
	ty := y - (h * originY)

	// Rotate
	options.GeoM.Translate(-w/2, -h/2)
	options.GeoM.Rotate(rotation * (2 * math.Pi))
	options.GeoM.Translate(w, h)

	// Position
	options.GeoM.Translate(tx, ty)

	screen.DrawImage(img, options)
}

func GetOriginPoint(posX, posY, rectWidth, rectHeight, rotation, originX, originY float64) (x, y float64) {
	// Local coordinates, rectangle is the screen basically
	localX := originX * rectWidth
	localY := originY * rectHeight

	// Move origin to the middle of the rectangle, so x and y go both to positive
	// and negative directions.
	centeredX := localX - (rectWidth / 2)
	centeredY := localY - (rectHeight / 2)

	// Ensure rotation is within [0;1) bounds, negate to follow clockwise rotation
	// direction.
	angle := -EuclideanMod(rotation, 1) * (2 * math.Pi)

	// Rotate the centered coordinates.
	rotatedCenteredX := centeredX*math.Cos(angle) + centeredY*math.Sin(angle)
	rotatedCenteredY := -centeredX*math.Sin(angle) + centeredY*math.Cos(angle)

	// Center of the rectangle on the plane.
	rectCenterX := posX + (rectWidth / 2)
	rectCenterY := posY + (rectHeight / 2)

	// Add up center of the rectangle on the plane with its centered
	// coordinates.
	x = rectCenterX + rotatedCenteredX
	y = rectCenterY + rotatedCenteredY

	return x, y
}
