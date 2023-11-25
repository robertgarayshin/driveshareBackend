package write

import (
	"driveshare_backend/internal/database/read"
	"driveshare_backend/internal/models"
)

func ChangeBookingStatus(status bool, id int) {
	for i := range read.CarDbMock {
		if read.CarDbMock[i].Id == id {
			read.CarDbMock[i].Booked = status
		}
	}
}

func PostCar(car models.Car) {
	read.CarDbMock = append(read.CarDbMock, car)
}

func DeleteCar(id int) {
	for i := range read.CarDbMock {
		if read.CarDbMock[i].Id == id {
			read.CarDbMock[i] = read.CarDbMock[len(read.CarDbMock)-1] // Copy last element to index i.
			read.CarDbMock[len(read.CarDbMock)-1] = models.Car{}      // Erase last element (write zero value).
			read.CarDbMock = read.CarDbMock[:len(read.CarDbMock)-1]   // Truncate slice.
		}
	}
}
