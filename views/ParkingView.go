// views/ParkingView.go
package views

import (
    "github.com/hajimehoshi/ebiten"
    "parkingLot/controllers"
)

type ParkingLotView struct {
    parkingLotController *controllers.ParkingLotController
}

func NewParkingLotView(plc *controllers.ParkingLotController) *ParkingLotView {
    return &ParkingLotView{
        parkingLotController: plc,
    }
}

func (plv *ParkingLotView) Update(screen *ebiten.Image) error {
    parkingLot := plv.parkingLotController.GetParkingLot()
    for i:= range parkingLot.Semaphore {
        opts := &ebiten.DrawImageOptions{}
        // Calculate x and y based on the vehicle's position.
        x := float64((i % 10) * 60 + 30)
        y := float64((i / 10) * 60 + 30)
        opts.GeoM.Translate(x, y)

        if i == 0 {  // Vehicle is leaving.
            screen.DrawImage(cOutImage, opts)
        } else {  // Vehicle is entering.
            screen.DrawImage(cImage, opts)
        }
    }

    return nil
}

