package service

import (
	"Trovero/Backend/internal/apperrors"
	"Trovero/Backend/internal/models"
	"Trovero/Backend/internal/repository"
	"errors"

	"gorm.io/gorm"
)

type addressService struct {
	addressRepository   repository.AddressRepository
	userService			UserService

}

type AddressService interface {
	CreateAddress(address *models.Address,UserID int64) error
	GetAll(UserID int64)([]models.Address,error)
	DeleteAddressByID(ID,UserID int64)error
}

func NewAddressService(addressRepository repository.AddressRepository,userService UserService) AddressService {
	return &addressService{
		addressRepository: addressRepository,
		userService:userService,
	}
}

func (s *addressService) CreateAddress(address *models.Address,UserID int64) error{
	user,err:=s.userService.GetUserByID(UserID)
	if err!=nil{
		if errors.Is(err,gorm.ErrRecordNotFound){
			return apperrors.ErrUserNotFound
		}
		return err
	}

	address.UserID=user.ID
	
	err=s.addressRepository.CreateAddress(address)
	if err!=nil{
		return err
	}
	return nil

}

func (s *addressService)GetAll(UserID int64)([]models.Address,error){
	return s.addressRepository.GetAll(UserID)

}

func (s *addressService)DeleteAddressByID(ID,UserID int64)error{
	
	err:=s.addressRepository.DeleteAddressByID(ID,UserID)
	if err!=nil{
		if errors.Is(err,gorm.ErrRecordNotFound){
			return apperrors.ErrAddressNotFound
		}
		return err
	}

	return nil

}
