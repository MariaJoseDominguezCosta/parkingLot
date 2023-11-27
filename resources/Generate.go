//go:generate go run github.com/hajimehoshi/file2byteslice/cmd/file2byteslice@v1.0.0 -package=images -input=./images/car.png -output=./images/Car.go  -var=CarPng
//go:generate go run github.com/hajimehoshi/file2byteslice/cmd/file2byteslice@v1.0.0 -package=images -input=./images/car_out.png -output=./images/CarOut.go -var=CarOutPng

//go:generate gofmt -s -w .

// Resources package contains font, image resources needed by the game
package resources
