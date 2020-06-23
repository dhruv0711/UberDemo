package uber

import (
	"github.com/jinzhu/gorm"
)

type Cabs struct {
	CabRegistrationNumber string `gorm:"unique;not null" json:"registration_number"`
	CabCurrentGps         string `gorm:"unique;not null" json:"current_gps"`
	CabDriverPhoneNumber  string `gorm:"unique;not null" json:"driver_phone"`
	CabDriverName         string `json:"driver_name"`
	CabMakeModel          string `json:"make_model"`
}

type Booking struct {
	gorm.Model
	CabRegistrationNumber string `json:"registration_number"`
	CabDriverName         string `json:"driver_name"`
	CabMakeModel          string `json:"make_model"`
	CabDriverPhoneNumber  string `json:"driver_phone"`
	CustomerName          string `json:"customer_name"`
	CustomerPhoneNumber   string `json:"customer_phone_number"`
	CustomerGps           string `json:"customer_gps"`
	PickUp                string `json:"pickup"`
	DropLocation          string `json:"dropoff"`
	IsActive              bool
}
