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
	GetAll(page,limit int)([]models.Coupon,error)
	GetCouponByID(id int64)(*models.Coupon,error)
	DeleteCouponByID(id int64)error
	UpdateCouponStatus(id int64,is_active bool)error
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


func (s *couponService)GetAll(page,limit int)([]models.Coupon,error){
	return s.couponRepository.GetAll(page,limit)
}


func (s *couponService)GetCouponByID(id int64)(*models.Coupon,error){

	coupon,err:=s.couponRepository.GetCouponByID(id)
	if err!=nil{
		if errors.Is(err,gorm.ErrRecordNotFound){
			return nil,apperrors.ErrCouponNotFound
		}
		return nil,err
	}

	return coupon,nil
}

func (s *couponService)DeleteCouponByID(id int64)error{

	err:=s.couponRepository.DeleteCouponByID(id)
	if err!=nil{
		if errors.Is(err,gorm.ErrRecordNotFound){
			return apperrors.ErrCouponNotFound
		}
		return err
	}

	return nil

}


func (s *couponService)UpdateCouponStatus(id int64,is_active bool)error{
	err:=s.couponRepository.UpdateCouponStatus(id,is_active)
	if err!=nil{
		if errors.Is(err,gorm.ErrRecordNotFound){
			return apperrors.ErrCouponNotFound
		}
		return err
	}
	return nil
}
