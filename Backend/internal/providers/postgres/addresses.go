package postgres

import (
	"Trovero/Backend/internal/models"
	"Trovero/Backend/internal/repository"

	"gorm.io/gorm"
)

type addressRepository struct {
	db *gorm.DB
}

func NewAddressRepository(db *gorm.DB) repository.AddressRepository {
	return &addressRepository{
		db: db,
	}
}

func (r *addressRepository) CreateAddress(address *models.Address) error{
	return r.db.Create(address).Error
}


func (r *addressRepository)GetAll(UserID int64)([]models.Address,error){
	var addresses []models.Address

	result:=r.db.Where("user_id=?",UserID).Find(&addresses)

	if result.Error!=nil{
		return []models.Address{},result.Error
	}

	return addresses,nil
}

func (r *addressRepository)DeleteAddressByID(ID,UserID int64)error{

	result:=r.db.Where("id=? AND user_id=?",ID,UserID).Delete(&models.Address{})

	if result.Error!=nil{
		return result.Error
	}

	if result.RowsAffected==0{
		return gorm.ErrRecordNotFound
	}

	return nil

}
