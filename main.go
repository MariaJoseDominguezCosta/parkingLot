package parkingLot

import (
	"parkingLot/controllers"
	"parkingLot/views"
)

func main() {

	views.CreateGUI()

	for {
		event := GetUserEvent()

		controllers.HandleEvent(event)
	}
}
