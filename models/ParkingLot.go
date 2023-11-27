package models

// ParkingLot struct represents the parking lot.
type ParkingLot struct {
	Capacity  int
	Semaphore chan int
}

// NewParkingLot creates a new parking lot with the given capacity.
func NewParkingLot(capacity int) *ParkingLot {
	return &ParkingLot{
		Capacity:  capacity,
		Semaphore: make(chan int, capacity),  // Buffered channel used as a semaphore.
	}
}
