package models

import (
	"fmt"
	"math/rand"
	"sync"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/storage"

	"time"
)

type Vehicle struct {
	Id              int
	ParkingDuration time.Duration // The amount of time the car will stay parked
	image           *canvas.Image
	space           int
	imageE          *canvas.Image
}

func NewVehicle(id int) *Vehicle {
	image := canvas.NewImageFromURI(storage.NewFileURI("./assets/car.png"))
	imageE := canvas.NewImageFromURI(storage.NewFileURI("./assets/car_out.png"))
	return &Vehicle{
		Id:              id,
		ParkingDuration: time.Duration(10+rand.Intn(20)) * time.Second,
		image:           image,
		imageE:          imageE,
		space:           0,
	}

}

func (c *Vehicle) GetId() int {
	return c.Id
}

func (v *Vehicle) GetVehicleImage() *canvas.Image {
	return v.image
}

func (v *Vehicle) EnterParking(p *ParkingLot, carsC *fyne.Container) {
	p.GetSpaces() <- v.GetId()
	p.GetMutex().Lock()

	spaces := p.GetParking()

	fmt.Printf("Vehicle %d has entered\n", v.GetId())

	for i := 0; i < 5; i++ {
		v.image.Move(fyne.NewPos(v.image.Position().X+20, v.image.Position().Y))
		time.Sleep(time.Millisecond * 200)
	}

	p.GetMutex().Unlock()

	for i := 0; i < len(spaces); i++ {

		if spaces[i] == false {

			spaces[i] = true
			v.space = i
			v.image.Move(fyne.NewPos(290, float32(15+(i+30))))
			break
		}

	}

	p.SetParking(spaces)
	carsC.Refresh()
}

func (v *Vehicle) ExitParking(p *ParkingLot, carsC *fyne.Container) {
	p.GetMutex().Lock()
	<-p.GetSpaces()

	spaces := p.GetParking()
	spaces[v.space] = false
	p.SetParking(spaces)

	fmt.Printf("Vehicle %d has left\n", v.GetId())
	p.GetMutex().Unlock()

	for i := 0; i < 5; i++ {
		v.imageE.Move(fyne.NewPos(v.imageE.Position().X-30, v.imageE.Position().Y))
		time.Sleep(time.Millisecond * 200)
	}

	carsC.Remove(v.imageE)
	carsC.Refresh()
}

func (v *Vehicle) Parking(p *ParkingLot, carsC *fyne.Container, wg *sync.WaitGroup) {
	for i := 0; i < 7; i++ {
		v.image.Move(fyne.NewPos(v.image.Position().X+20, v.image.Position().Y))
		time.Sleep(time.Millisecond * 200)
	}

	v.EnterParking(p, carsC)
	time.Sleep(v.ParkingDuration)

	carsC.Remove(v.image)
	v.imageE.Resize(fyne.NewSize(50, 30))
	p.Exit(carsC, v.imageE)
	v.ExitParking(p, carsC)
	wg.Done()
}
