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


func (r *couponRepository)GetAll(page,limit int)([]models.Coupon,error){
	var coupons []models.Coupon

	offset:=(page-1)*limit

	result:=r.db.Limit(limit).Offset(offset).Find(&coupons)
	if result.Error!=nil{
		return []models.Coupon{},result.Error
	}
	return coupons,nil
}


func (r *couponRepository)GetCouponByID(id int64)(*models.Coupon,error){
	var coupon models.Coupon

	result:=r.db.Where("id=?",id).First(&coupon)
	if result.Error!=nil{
		return nil,result.Error
	}
	return &coupon,nil
}


func (r *couponRepository)DeleteCouponByID(id int64)error{

	result:=r.db.Where("id=?",id).Delete(&models.Coupon{})
	if result.Error!=nil{
		return result.Error
	}

	if result.RowsAffected==0{
		return gorm.ErrRecordNotFound
	}
	return nil
}


func (r *couponRepository)UpdateCouponStatus(id int64,is_active bool)error{

	result:=r.db.Model(&models.Coupon{}).Where("id=?",id).Update("is_active",is_active)

	if result.Error!=nil{
		return result.Error
	}

	if result.RowsAffected==0{
		return gorm.ErrRecordNotFound
	}

	return nil

}

