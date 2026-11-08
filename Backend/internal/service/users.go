package service

import (
	"Trovero/Backend/internal/apperrors"
	"Trovero/Backend/internal/models"
	"Trovero/Backend/internal/repository"
	"Trovero/Backend/internal/utils"
	"Trovero/Backend/tokens"
	"errors"
	"fmt"
	"time"

	"gorm.io/gorm"
)

type userService struct {
	UserRepository  repository.UserRepository
	tokenMaker		tokens.Maker
	tokenDuration	time.Duration
}

type UserService interface{
	CreateUser(user *models.User)error
	Login(email,password string)(string,error)
	GetUserByID(UserID int64)(*models.User,error)
	GetAll()([]models.User,error)
	DeleteUserByID(UserID int64)error
}

func NewUserService (UserRepository repository.UserRepository,tokenMaker tokens.Maker,tokenDuration time.Duration)UserService{
	return &userService{
		UserRepository: UserRepository,
		tokenMaker: tokenMaker,
		tokenDuration: tokenDuration,
	}
}

func(s *userService)CreateUser(user *models.User)error{
	existing,err:=s.UserRepository.GetUserByEmail(user.Email)
	if err!=nil{
		if !errors.Is(err,gorm.ErrRecordNotFound){
			return err
		}
	}

	if existing!=nil{
		return apperrors.ErrUserAlreadyExists
	}

	hashedpassword,err:=utils.HashPassword(user.Password)
	if err!=nil{
		return err
	}

	user.Password=hashedpassword

	return s.UserRepository.CreateUser(user)
}

func(s *userService)Login(email,password string)(string,error){
	exisiting,err:=s.UserRepository.GetUserByEmail(email)
	if err!=nil{
		if errors.Is(err,gorm.ErrRecordNotFound){
			return "",apperrors.ErrInvalidCredentials
		}
		return "",err
	}
	
	err=utils.VerifyPassword(exisiting.Password,password)
	if err!=nil{
		return "",apperrors.ErrInvalidCredentials
	}

	accessToken,err:=s.tokenMaker.CreateToken(exisiting.ID,exisiting.Role,s.tokenDuration)
	if err!=nil{
		return "",fmt.Errorf("failed to generate tokens:%v",err)
	}
	return accessToken,nil
}

func(s *userService)GetUserByID(UserID int64)(*models.User,error){
	user,err:=s.UserRepository.GetUserByID(UserID)
	if err!=nil{
		if errors.Is(err,gorm.ErrRecordNotFound){
			return nil,apperrors.ErrUserNotFound
		}
		return nil,err
	}

	return user,nil
}

func(s *userService)DeleteUserByID(UserID int64)error{

	err:=s.UserRepository.DeleteUserByID(UserID)
	if err!=nil{
		if errors.Is(err,gorm.ErrRecordNotFound){
			return apperrors.ErrUserNotFound
		}
		return err
	}

	return nil
}


func(s *userService)GetAll()([]models.User,error){
	return s.UserRepository.GetAll()
}
