package models

import "time"

type Address struct {
	ID              int64  		`gorm:"primarykey" json:"id"`
	UserID          int64  		`gorm:"not null" json:"user_id"`
	City            string 		`gorm:"not null" json:"city"`
	Area            string 		`gorm:"not null" json:"area"`
	Street          string 		`gorm:"not null" json:"street"`
	BuildingNumber  int    		`gorm:"not null" json:"building_number"`
	ApartmentNumber *int    	`json:"apartment_number"`
	PhoneNumber     string 		`gorm:"not null" json:"phone_number"`
	CreatedAt       time.Time	`gorm:"autoCreateTime" json:"-"`
}	
