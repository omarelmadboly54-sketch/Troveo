package postgres

import (
	"Trovero/Backend/internal/models"
	"Trovero/Backend/internal/repository"

	"gorm.io/gorm"
)

type couponRepository struct {
	db *gorm.DB
}

func NewCouponRepository(db *gorm.DB) repository.CouponRepository {
	return &couponRepository{
		db: db,
	}
}

func (r *couponRepository) CreateCoupon(coupon *models.Coupon) error{
	return r.db.Create(coupon).Error
}


