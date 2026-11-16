package postgres

import (
	"Trovero/Backend/internal/models"
	"Trovero/Backend/internal/repository"

	"gorm.io/gorm"
)

type usercouponRepository struct {
	db *gorm.DB
}


func NewUserCouponRepository (db *gorm.DB)repository.UserCouponRepository{
	return &usercouponRepository{
		db: db,
	}
}

func(r *usercouponRepository)AssignCouponToUser(usercoupon *models.UserCoupon)error{
	return r.db.Create(usercoupon).Error
}


func(r *usercouponRepository)GetMyCoupons(UserID int64,page,limit int)([]models.UserCoupon,error){
	var userCoupons []models.UserCoupon

	offset:=(page-1)*limit

	result:=r.db.Where("user_id=?",UserID).Limit(limit).Offset(offset).Find(&userCoupons)
	if result.Error!=nil{
		return []models.UserCoupon{},result.Error
	}
	
	return userCoupons,nil
}


func(r *usercouponRepository)GetMyCoupon(id,UserID int64)(*models.UserCoupon,error){

	var userCoupon models.UserCoupon

	result:=r.db.Where("id=? AND user_id=?",id,UserID).First(&userCoupon)

	if result.Error!=nil{
		return nil,result.Error
	}
	
	return &userCoupon,nil
}