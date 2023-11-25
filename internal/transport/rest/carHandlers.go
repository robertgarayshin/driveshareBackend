package rest

import (
	"driveshare_backend/internal/database/read"
	"driveshare_backend/internal/database/write"
	"driveshare_backend/internal/models"
	"encoding/json"
	"github.com/gorilla/mux"
	"net/http"
	"strconv"
)

// PostCarHandler
//
//	@Summary		New car publishing endpoint
//	@ID				postCar
//	@Accept json
//	@Tags car
//	@Security		ApiKeyAuth
//	@Param Car body models.Car true "query params"
//	@Success		200	{string}	string	"ok"
//	@Router			/car/new [post]
func PostCarHandler(w http.ResponseWriter, r *http.Request) {
	decoder := json.NewDecoder(r.Body)
	var car models.Car
	err := decoder.Decode(&car)
	if err != nil {
		return
	}
	write.PostCar(car)
}

// GetCarByIdHandler
//
//	@Summary		Get car by id
//	@ID				getCarById
//	@Tags 			car
//	@Security		ApiKeyAuth
//	@Param       	id   path      int  true  "Car ID"
//	@Success		200	{object}	models.Car	"ok"
//	@Router			/car/{id} [get]
func GetCarByIdHandler(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id, _ := strconv.Atoi(params["id"])
	car, err := json.Marshal(read.GetCarById(id))
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.Write(car)
}

// DeleteCarHandler
//
//	@Summary		Delete car by id
//	@ID				deleteCarById
//	@Tags 			car
//	@Security		ApiKeyAuth
//	@Param       	id   path      int  true  "Car ID"
//	@Success		200	{string}	string	"ok"
//	@Router			/car/{id} [delete]
func DeleteCarHandler(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id, _ := strconv.Atoi(params["id"])
	write.DeleteCar(id)
}

// GetAllCarsHandler
//
//	@Summary		Get all cars
//	@ID				getAllCars
//	@Tags 			car
//	@Security		ApiKeyAuth
//	@Success		200	{array}	models.Car	"ok"
//	@Router			/car [get]
func GetAllCarsHandler(w http.ResponseWriter, r *http.Request) {
	cars, err := json.Marshal(read.GetAllCars())
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.Write(cars)
}

// BookCarHandler
//
//	@Summary		Book car by id
//	@ID				bookCarById
//	@Tags 			car
//	@Security		ApiKeyAuth
//	@Param       	id   path      int  true  "Car ID"
//	@Success		200	{string}	string	"ok"
//	@Router			/car/book/{id} [put]
func BookCarHandler(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id, _ := strconv.Atoi(params["id"])
	write.ChangeBookingStatus(true, id)
}

// CancelBookHandler
//
//	@Summary		Cancel booking
//	@ID				cancelBooking
//	@Tags 			car
//	@Security		ApiKeyAuth
//	@Param       	id   path      int  true  "Car ID"
//	@Success		200	{string}	string	"ok"
//	@Router			/car/cancel/{id} [put]
func CancelBookHandler(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id, _ := strconv.Atoi(params["id"])
	write.ChangeBookingStatus(false, id)
}

// GetAvailableCarsHandler
//
//	@Summary		Get available cars
//	@ID				getAvailableCars
//	@Tags 			car
//	@Security		ApiKeyAuth
//	@Success		200	{array}	models.Car	"ok"
//	@Router			/car/available [get]
func GetAvailableCarsHandler(w http.ResponseWriter, r *http.Request) {
	cars, err := json.Marshal(read.GetAllAvailableCars())
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.Write(cars)
}
