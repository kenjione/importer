package model

import (
	v "github.com/asaskevich/govalidator"
)

type Location struct {
	ID           int64  `json:"id"`
	IPaddress    string `json:"ip_address" sql:"ip_address,unique:idx_locations_on_ip_address" valid:"ip"`
	CountryCode  string `json:"country_code" valid:"stringlength(2|2)"`
	Country      string `json:"country" valid:"required"`
	City         string `json:"city" valid:"required"`
	Latitude     string `json:"latitude" valid:"latitude"`
	Longitude    string `json:"longitude" valid:"longitude"`
	MysteryValue string `json:"mystery_value" valid:"required"`
}

func (l *Location) Validate() error {
	_, err := v.ValidateStruct(l)
	return err
}
