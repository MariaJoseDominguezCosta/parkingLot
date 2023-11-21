package models

import ("sync"
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
)
const SPOTS = 20

type ParkingLot struct {
	Spaces chan int
	// Vehicles []Vehicle
	// Sem      chan struct{}
	Parking      [SPOTS]bool
	mutex      *sync.Mutex
	// inputQueue chan *Vehicle
	// Spots    [SPOTS]int
}

func NewParkingLot(capacity chan int, m *sync.Mutex) *ParkingLot {
	return &ParkingLot{
		Spaces: capacity,
		mutex:    m,
		Parking:    [SPOTS]bool{},
	}
}

func (p *ParkingLot) GetSpaces() chan int {
	return p.Spaces
}

func (p *ParkingLot) GetMutex() *sync.Mutex {
	return p.mutex
}

func (p *ParkingLot) GetParking() [SPOTS]bool {
	return p.Parking
}

func (p *ParkingLot) SetParking(parking [SPOTS]bool) {
	p.Parking = parking
}

func (p *ParkingLot) Exit(carsC *fyne.Container, carI *canvas.Image) {
	carI.Move(fyne.NewPos(205, 350))
	carsC.Add(carI)
	carsC.Refresh()
}