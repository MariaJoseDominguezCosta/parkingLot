package models

type Vehicle struct {}

func NewVehicle() *Vehicle {
	return &Vehicle{}
}

func (v *Vehicle) EnterParking(){
	Parking.Spaces.Wait()

	Parking.Mutex.Wait()

	for i := 0; i < len(Parking.Spots); i++ {
		if Parking.Spots[i].TryWait(){

			TakeSpot(i)
			break
		}
	}

	Parking.Mutex.Signal()
	Parking.Space.Signal()
	
}