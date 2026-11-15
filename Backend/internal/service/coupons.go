package service

import (
	"Trovero/Backend/internal/apperrors"
	"Trovero/Backend/internal/models"
	"Trovero/Backend/internal/repository"
	"errors"
	"time"

	"gorm.io/gorm"
)

type couponService struct {
	couponRepository repository.CouponRepository
}

type CouponService interface {
	CreateCoupon(coupon *models.Coupon) error
}

func NewCouponService(couponRepository repository.CouponRepository) CouponService {
	return &couponService{
		couponRepository: couponRepository,
	}
}

func (s *couponService) CreateCoupon(coupon *models.Coupon) error {
	if coupon.DiscountType == "percentage" && coupon.Value > 100 {
		return apperrors.ErrInvalidCouponValue
	}

	if !coupon.ExpiredAt.After(time.Now()) {
		return apperrors.ErrInvalidCouponExpiration
	}

	err:=s.couponRepository.CreateCoupon(coupon)
	if err!=nil{
		if errors.Is(err,gorm.ErrDuplicatedKey){
			return apperrors.ErrCouponCodeExists
		}
		return err
	}
	return nil
}
