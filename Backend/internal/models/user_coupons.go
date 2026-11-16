package models

import (
	"time"
)

type UserCoupon struct {
	ID         int64 `gorm:"primarykey" json:"id"`
	UserID     int64 `gorm:"not null" json:"user_id"`
	CouponID   int64 `gorm:"not null" json:"coupon_id"`
	IsUsed     bool	 `gorm:"not null;default:false" json:"is_used"`	 
	AssignedAt time.Time `gorm:"autoCreateTime" json:"-"`
}
