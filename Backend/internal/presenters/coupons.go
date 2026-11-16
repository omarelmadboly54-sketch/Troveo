package presenters

import (
	"Trovero/Backend/internal/models"
	"time"
)

type CreateCouponReq struct {
	Code               string  		`json:"code" binding:"required"`
	DiscountType       string  		`json:"discount_type" binding:"required,oneof=percentage fixed"`
	Value              float64 		`json:"value" binding:"required,gt=0"`
	MinimumOrderAmount *float64 	`json:"minimum_order_amount" binding:"required,gte=0"`
	ExpiredAt          time.Time	`json:"expired_at" binding:"required"`
	IsActive		   *bool		`json:"is_active" binding:"required" `  		
}

type UpdateCouponStatusReq struct{
	IsActive			*bool		`json:"is_active" binding:"required"`		
}

func(c *CreateCouponReq) ToCoupon()models.Coupon{
	return models.Coupon{
		Code: c.Code,
		DiscountType: c.DiscountType,
		Value: c.Value,
		MinimumOrderAmount: *c.MinimumOrderAmount,
		ExpiredAt: c.ExpiredAt,
		IsActive: *c.IsActive,
	}
}
