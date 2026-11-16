package repository

import "Trovero/Backend/internal/models"

type UserCouponRepository interface {
	AssignCouponToUser(usercoupon *models.UserCoupon) error
	GetMyCoupons(UserID int64,page,limit int)([]models.UserCoupon,error)
	GetMyCoupon(id,UserID int64)(*models.UserCoupon,error)
}




