package models

import "time"

type Review struct {
	ID        int64  	`gorm:"primarykey" json:"id"`
	ProductID int64  	`gorm:"not null" json:"product_id"`
	Comment   string 	`gorm:"not null" json:"comment"`
	Rating    int    	`gorm:"not null" json:"rating"`
	UserID    int64  	`gorm:"not null" json:"user_id"`
	CreatedAt time.Time	`gorm:"autoCreateTime" json:"-"`
}
