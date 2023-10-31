package models

import "sync"

type Parking struct {
	Mutex sync.S
	Spaces sync.Semaphore
	Spots []*sync.Semaphore
}

var( Parking = &Parking{
	Mutex: sync.NewSemaphore(1),
	Spaces: sync.NewSemaphore(20),
	Spots: make([]*sync.Semaphore, 20),
}
)
for i := 0; i < 20; i++ {
	Parking.Spots[i] = sync.NewSemaphore(1)
}
