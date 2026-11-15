package models

import "time"

type Coupon struct {
	ID                 int64   		`gorm:"primarykey" json:"id"`
	Code               string  		`gorm:"unique;not null" json:"code"`
	DiscountType       string  		`gorm:"not null" json:"discount_type"`
	Value              float64 		`gorm:"type:numeric(10,2);not null" json:"value"`
	MinimumOrderAmount float64 		`gorm:"type:numeric(10,2);not null" json:"minimum_order_amount"`
	ExpiredAt          time.Time	`gorm:"not null" json:"expired_at"`
	IsActive		   bool			`gorm:"not null;default:true" json:"is_active"`
	CreatedAt		   time.Time	`gorm:"autoCreateTime" json:"-"`	
}
