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

// RoundTo rounds the value to a specific precision.
func RoundTo(value float64, precision int) float64 {
	factor := math.Pow(10, float64(precision))
	return math.Round(value*factor) / factor
}

// TransformedLocalPoint calculates where local point will end up after
// transformations are applied to the rectangle, such as rotation. Origin of
// transformation is specified as originX and originY in range [0;1].
func TransformedLocalPoint(localX, localY, rectWidth, rectHeight, rotation, originX, originY float64) (x, y float64) {
	// Set center (0,0) to the origin point
	centeredX := localX - (rectWidth * originX)
	centeredY := localY - (rectHeight * originY)

	// Negate for clockwise rotation
	angle := -(rotation * 2 * math.Pi)

	// Apply rotation transformation
	rotatedX := centeredX*math.Cos(angle) + centeredY*math.Sin(angle)
	rotatedY := -centeredX*math.Sin(angle) + centeredY*math.Cos(angle)

	// Return point (0,0) to the top-left corner
	x = RoundTo(rotatedX+(rectWidth*originX), 5)
	y = RoundTo(rotatedY+(rectHeight*originY), 5)

	return x, y
}

// DirectionMultipliers calculates direction multipliers for each coordinate.
// Multipliers will always add up to 1. If direction is along a single axis,
// then multiplier for that axis will be 1, and 0 for the other axis. If
// direction is perfectly diagonal, meaning it moves equally along each axis,
// the multipliers will be 0.5 for both.
//
// Direction is in range [0;1) with 0 being north, 0.25 - east, 0.5 - south and
// 0.75 - west. Note: the function still works as expected with direction
// outside the range, including negative values.
//
// Triangular wave functions are used to calculate the multipliers linearly:
//
// x(r)=4|(r-0.25)%1-0.5|-1
//
// y(r)=4|(r-0.5)%1-0.5|-1
func DirectionMultipliers(direction float64) (mx, my float64) {
	mx = (4*math.Abs(EuclideanMod(direction-0.25, 1)-0.5) - 1)
	my = (4*math.Abs(EuclideanMod(direction-0.5, 1)-0.5) - 1)
	return mx, my
}

// DrawImage is similar to ebiten.DrawImage, but allows specifying origin point
// in [0;1] range and rotation in full rotations.
func DrawImage(screen *ebiten.Image, img *ebiten.Image, posX, posY, rotation, originX, originY float64) {
	options := &ebiten.DrawImageOptions{
		Filter: ebiten.FilterLinear,
	}

	w := float64(img.Bounds().Dx())
	h := float64(img.Bounds().Dy())

	// Set center (0,0) to the origin point
	options.GeoM.Translate(-originX*w, -originY*h)

	// Apply rotation
	options.GeoM.Rotate(rotation * (2 * math.Pi))

	// Move to the position on the screen
	options.GeoM.Translate(posX, posY)

	// Return point (0,0) to the top-left corner
	options.GeoM.Translate(originX*w, originY*h)

	// Set center (0,0) to the center of the image
	options.GeoM.Translate(-w/2, -h/2)

	screen.DrawImage(img, options)
}

func ImageSize(img *ebiten.Image) (width, height float64) {
	width = float64(img.Bounds().Dx())
	height = float64(img.Bounds().Dy())
	return width, height
}
