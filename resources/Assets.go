// parkingLot/resources/assets.go
package resources

import (
	"bytes"
	"github.com/hajimehoshi/ebiten"
	"image/png"
	"parkingLot/resources/images"
)

type Cars struct {
	CarIn  *ebiten.Image
	CarOut *ebiten.Image
}

type Assets struct {
	Cars *Cars
}

// LoadAssets converts the character images(png, jpg, ...) to ebiten image format and loads fonts.
func LoadAssets() (*Assets, error) {

	cars, carsErr := loadCars()
	if carsErr != nil {
		return nil, carsErr
	}

	return &Assets{
		Cars: cars,
	}, nil
}

func loadCars() (*Cars, error) {
	cImage, cImageErr := png.Decode(bytes.NewReader(images.CarPng))
	if cImageErr != nil {
		return nil, cImageErr
	}

	cOutImage, cOutImageErr := png.Decode(bytes.NewReader(images.CarOutPng))
	if cOutImage != nil {
		return nil, cOutImageErr
	}

	carOut, carOutErr := ebiten.NewImageFromImage(cOutImage, ebiten.FilterDefault)
	if carOutErr != nil {
		return nil, carOutErr
	}

	carIn, carInErr := ebiten.NewImageFromImage(cImage, ebiten.FilterDefault)
	if carInErr != nil {
		return nil, carInErr
	}

	return &Cars{
		CarIn:  carIn,
		CarOut: carOut,
	}, nil
}
