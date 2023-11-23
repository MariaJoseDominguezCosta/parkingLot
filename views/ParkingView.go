// views/ParkingView.go
package views

import (
	"fmt"
	"parkingLot/controllers"
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

type ParkingView struct {
	App              fyne.App
	Window           fyne.Window
	ParkingController *controllers.ParkingController
	Spaces           []*canvas.Rectangle
	Vehicles         []*canvas.Image
}

func NewParkingView(pc *controllers.ParkingController) *ParkingView {
	a := app.New()
	w := a.NewWindow("Parking Lot")

	return &ParkingView{
		App:               a,
		Window:            w,
		ParkingController: pc,
		Spaces:            make([]*canvas.Rectangle, len(pc.ParkingLot.Spaces)),
		Vehicles:          make([]*canvas.Image, 100), // Assuming a maximum of 100 vehicles
	}
}

func (pv *ParkingView) Show() error {
	c := container.NewVBox()

	for _, v := range pv.ParkingController.ParkingLot.Spaces {
		if v != nil {
			c.Add(v.Image)
			c.Add(widget.NewLabel(fmt.Sprintf("Vehicle %d", v.Id)))
		} else {

			c.Add(widget.NewLabel("Empty Space"))
		}
	}

	pv.Window.SetContent(c)
	pv.Window.ShowAndRun()

	return nil
}

func (pv *ParkingView) Update() {
	c := container.NewVBox()

	for _, v := range pv.ParkingController.ParkingLot.Spaces {
		if v != nil {
			c.Add(v.Image)
			c.Add(widget.NewLabel(fmt.Sprintf("Vehicle %d", v.Id)))
		} else {
			c.Add(widget.NewLabel("Empty Space"))
		}
	}

	pv.Window.SetContent(c)
	pv.Window.Refresh()
}