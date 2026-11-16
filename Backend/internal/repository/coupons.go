package repository

import "Trovero/Backend/internal/models"

type CouponRepository interface {
	CreateCoupon(coupon *models.Coupon) error
	GetAll(page,limit int)([]models.Coupon,error)
	GetCouponByID(id int64)(*models.Coupon,error)
	DeleteCouponByID(id int64)error
	UpdateCouponStatus(id int64,is_active bool)error
}




