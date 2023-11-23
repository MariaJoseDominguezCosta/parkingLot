package models

import "errors"

type ParkingSpace int
type ParkingLot struct {
	Spaces     []*Vehicle
	AvailSpots []int // Queue to store the indices of available parking spaces
}

func NewParkingLot(capacity ParkingSpace) (*ParkingLot) {
	spaces := make([]*Vehicle, capacity) // Initialize the slice with nil values
	availSpots := make([]int, capacity)

	for i := 0; i < int(capacity); i++ {
	    availSpots[i] = i
	}

	return &ParkingLot{
		Spaces:     spaces,
		AvailSpots: availSpots,
	}
}

func (p *ParkingLot) Park(v *Vehicle) error {
	if len(p.AvailSpots) == 0 {
		return errors.New("parking lot is full")
	}

	// Get the first available spot from the queue
	spotIndex := p.AvailSpots[0]
	p.AvailSpots = p.AvailSpots[1:]

	// Park the vehicle in the selected spot
	p.Spaces[spotIndex] = v

	return nil
}