package repository

import "Trovero/Backend/internal/models"

type AddressRepository interface {
	CreateAddress(address *models.Address) error
	GetAll(UserID int64)([]models.Address,error)
	DeleteAddressByID(ID,UserID int64)error
}