package service

import (
	"Trovero/Backend/internal/apperrors"
	"Trovero/Backend/internal/models"
	"Trovero/Backend/internal/repository"
	"errors"
	"time"

	"gorm.io/gorm"
)

type usercouponService struct {
	usercouponRepository repository.UserCouponRepository
	userService          UserService
	couponService        CouponService
}

type UserCouponService interface {
	AssignCouponToUser(userCoupon *models.UserCoupon) error
	GetMyCoupons(UserID int64, page, limit int) ([]models.UserCoupon, error)
	GetMyCoupon(id,UserID int64)(*models.UserCoupon,error)
}

func NewUserCouponService(usercouponRepository repository.UserCouponRepository, userService UserService, couponService CouponService) UserCouponService {
	return &usercouponService{
		usercouponRepository: usercouponRepository,
		userService:          userService,
		couponService:        couponService,
	}
}

func (s *usercouponService) AssignCouponToUser(userCoupon *models.UserCoupon) error {
	user, err := s.userService.GetUserByID(userCoupon.UserID)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return apperrors.ErrUserNotFound
		}
		return err
	}

	if user.AccountType != "customer" {
		return apperrors.ErrCouponAssignmentToNonCustomer
	}

	coupon, err := s.couponService.GetCouponByID(userCoupon.CouponID)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return apperrors.ErrCouponNotFound
		}
		return err
	}

	if !coupon.IsActive {
		return apperrors.ErrCouponNotActive
	}

	if !coupon.ExpiredAt.After(time.Now()) {
		return apperrors.ErrCouponExpired
	}

	err = s.usercouponRepository.AssignCouponToUser(userCoupon)
	if err != nil {
		if errors.Is(err, gorm.ErrDuplicatedKey) {
			return apperrors.ErrCouponAlreadyAssigned
		}
		return err
	}

	return nil

}

func (s *usercouponService) GetMyCoupons(UserID int64, page, limit int) ([]models.UserCoupon, error) {
	return s.usercouponRepository.GetMyCoupons(UserID, page, limit)
}


func (s *usercouponService)GetMyCoupon(id,UserID int64)(*models.UserCoupon,error){

	userCoupon,err:=s.usercouponRepository.GetMyCoupon(id,UserID)
	if err!=nil{
		if errors.Is(err,gorm.ErrRecordNotFound){
			return nil,apperrors.ErrCouponNotFound
		}
		return nil,err
	}

	return userCoupon,nil
}
