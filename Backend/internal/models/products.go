package models

import "time"

type Product struct {
	ID          	int64     `gorm:"primarykey" json:"id"`
	Name        	string    `gorm:"not null" json:"name"`
	Description 	string    `gorm:"not null" json:"description"`
	Price       	float64   `gorm:"type:numeric(10,2);not null" json:"price"`
	Photo       	string    `gorm:"not null" json:"photo"`
	PhotoPublicID	string	  `json:"-"`	
	Quantity    	int       `gorm:"not null;default:0" json:"quantity"`
	CategoryID  	int64     `gorm:"not null" json:"category_id"`
	SellerID    	int64     `gorm:"not null" json:"seller_id"`
	CreatedAt   	time.Time `gorm:"autoCreateTime" json:"-"`
}

