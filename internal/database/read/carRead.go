package read

import (
	"driveshare_backend/internal/models"
	"fmt"
)

var CarDbMock = []models.Car{
	{
		Id:       1,
		Brand:    "Chery",
		Model:    "Omoda C5",
		Produced: "2022",
		Price:    3398,
		Booked:   false,
	},
	{
		Id:       2,
		Brand:    "Mini",
		Model:    "Coupe",
		Produced: "2014",
		Price:    4392,
		Booked:   false,
	},
	{
		Id:       3,
		Brand:    "Mercedes-Benz",
		Model:    "E-Class",
		Produced: "2014",
		Price:    4792,
		Booked:   false,
	},
}

func GetAllCars() []models.Car {
	return CarDbMock
}

func GetCarById(id int) models.Car {
	for i := range CarDbMock {
		if CarDbMock[i].Id == id {
			return CarDbMock[i]
		}
	}
	return models.Car{}
}

// GetAllAvailableCars TODO: проверить
func GetAllAvailableCars() []models.Car {
	res := make([]models.Car, 0)
	for _, i := range CarDbMock {
		if i.Booked == false {
			res = append(res, i)
			fmt.Println(res)
		}
	}
	return res
}
