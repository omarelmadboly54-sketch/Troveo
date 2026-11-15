package repository

import "Trovero/Backend/internal/models"

type UserRepository interface {
	CreateUser(user *models.User)error
	GetUserByEmail(email string)(*models.User,error)
	GetAll(page int,limit int,accountType string)([]models.User,error)
	GetUserByID(UserID int64)(*models.User,error)
	DeleteUserByID(UserID int64)error
	UpdateUserStatus(UserID int64,IsActive bool)error
}


