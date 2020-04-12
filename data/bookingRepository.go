package data

import (
	"github.com/rbrto/bookings/models"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

// metodo de BookingRepository (Create, GetAll, Delete)
type BookingRepository struct {
	C *mgo.Collection
}

// metodo de BookingRepository se envia como parametros (booking *models.Booking)
func (r *BookingRepository) Create(booking *models.Booking) error {
	obj_id := bson.NewObjectId()
	booking.Id = obj_id
	err := r.C.Insert(&booking)
	return err
}

// metodo de BookingRepository devuelve arreglo []models.Booking no tiene parametros
func (r *BookingRepository) GetAll() []models.Booking {
	var bookings []models.Booking
	iter := r.C.Find(nil).Iter()
	result := models.Booking{}
	for iter.Next(&result) {
		bookings = append(bookings, result)
	}
	return bookings
}

func (r *BookingRepository) Delete(id string) error {
	err := r.C.Remove(bson.M{"_id": bson.ObjectIdHex(id)})
	return err
}
