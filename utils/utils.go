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

func DrawImage(screen *ebiten.Image, img *ebiten.Image, x, y, originX, originY, rotation float64) {
	options := &ebiten.DrawImageOptions{}

	imgWidth := float64(img.Bounds().Dx())
	imgHeight := float64(img.Bounds().Dy())

	x -= imgWidth * originX
	y -= imgHeight * originY

	// Rotate
	options.GeoM.Translate(-imgWidth/2, -imgHeight/2)
	options.GeoM.Rotate(rotation * (2 * math.Pi))
	options.GeoM.Translate(imgWidth, imgHeight)

	// Position
	options.GeoM.Translate(x, y)

	screen.DrawImage(img, options)
}
