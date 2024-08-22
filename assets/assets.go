package assets

import (
	"embed"
	"fmt"
	"image"
	"io/fs"
	"reflect"

	_ "image/png"

	"github.com/hajimehoshi/ebiten/v2"
)

//go:embed images/*
var images embed.FS

type Assets struct {
	Player    Player
	Asteroids []*ebiten.Image `path:"images/asteroids/*.png"`
}

type Player struct {
	Ship  *ebiten.Image `path:"images/player/PlayerShip.png"`
	Laser *ebiten.Image `path:"images/player/PlayerLaser.png"`
}

func LoadAssets() (*Assets, error) {
	var a Assets

	err := populateAssets(reflect.ValueOf(&a).Elem())
	if err != nil {
		return nil, fmt.Errorf("error populating assets: %v", err)
	}

	return &a, nil
}

func populateAssets(ref reflect.Value) error {
	for i := 0; i < ref.NumField(); i++ {
		field := ref.Field(i)
		path := ref.Type().Field(i).Tag.Get("path")

		switch field.Type().Kind() {
		case reflect.Struct:
			err := populateAssets(field)
			if err != nil {
				return err
			}
		case reflect.Slice:
			images, err := loadEbitenImages(path)
			if err != nil {
				return fmt.Errorf("error loading images: %v", err)
			}
			field.Set(reflect.ValueOf(images))
		default:
			img, err := loadEbitenImage(path)
			if err != nil {
				return fmt.Errorf("error loading an image: %v", err)
			}
			field.Set(reflect.ValueOf(img))
		}
	}

	return nil
}

func loadEbitenImages(path string) ([]*ebiten.Image, error) {
	matches, err := fs.Glob(images, path)
	if err != nil {
		return nil, fmt.Errorf("error finding files matching %q: %v", path, err)
	}

	images := make([]*ebiten.Image, 0, len(matches))
	for _, match := range matches {
		img, err := loadEbitenImage(match)
		if err != nil {
			return nil, fmt.Errorf("error loading image: %v", err)
		}
		images = append(images, img)
	}

	return images, nil
}

func loadEbitenImage(path string) (*ebiten.Image, error) {
	file, err := images.Open(path)
	if err != nil {
		return nil, fmt.Errorf("error opening path %q: %v", path, err)
	}

	decodedImg, _, err := image.Decode(file)
	if err != nil {
		return nil, fmt.Errorf("error decoding image %q: %v", path, err)
	}

	img := ebiten.NewImageFromImage(decodedImg)

	return img, nil
}
