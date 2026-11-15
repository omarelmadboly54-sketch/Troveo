package repository

import "Trovero/Backend/internal/models"

type CouponRepository interface {
	CreateCoupon(coupon *models.Coupon) error
}