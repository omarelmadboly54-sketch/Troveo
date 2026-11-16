package presenters

import "Trovero/Backend/internal/models"

type AssignCouponToUserReq struct {
	UserID   int64 `json:"user_id" binding:"required"`
	CouponID int64 `json:"coupon_id" binding:"required"`
}

func (a *AssignCouponToUserReq) ToUserCoupon() models.UserCoupon{
	return models.UserCoupon{
		UserID: a.UserID,
		CouponID: a.CouponID,
	}
}
