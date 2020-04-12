package controllers

import (
	"encoding/json"
	"net/http"

	"github.com/rbrto/bookings/common"
	"github.com/rbrto/bookings/data"
)

// Handler for HTTP Post - "/bookins"
// Create a new Booking document
func CreateBooking(w http.ResponseWriter, r *http.Request) {
	var dataResource BookingResource
	// Decode the incoming Booking json to Booking struct go format is saved into dataResource
	err := json.NewDecoder(r.Body).Decode(&dataResource)
	if err != nil {
		common.DisplayAppError(w, err, "Invalid Booking data", 500)
		return
	}
	booking := &dataResource.Data
	// Create new context associated to Mongo for each request
	context := NewContext()
	defer context.Close()
	c := context.DbCollection("bookings") // c =  *mgo.Collection bookings

	// Create Booking
	// Create and initialize C mgo.Collection variable from struct BookingRepository
	repo := &data.BookingRepository{c}
	repo.Create(booking)
	// Create response data (Marshall convert to json)
	j, err := json.Marshal(dataResource)
	if err != nil {
		common.DisplayAppError(w, err, "An unexpected error has occurred", 500)
		return
	}

	// Send response back
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(j)
}

func GetBookings(w http.ResponseWriter, r *http.Request) {
	// Create new context associated to Mongo for each request
	context := NewContext()
	defer context.Close()
	c := context.DbCollection("bookings")
	repo := &data.BookingRepository{c}
	// Get all bookings
	bookings := repo.GetAll()
	// Create response data json
	j, err := json.Marshal(BookingsResource{Data: bookings})
	if err != nil {
		common.DisplayAppError(w, err, "An unexpected error has occurred", 500)
		return
	}
	// Send response back
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(j)
}
