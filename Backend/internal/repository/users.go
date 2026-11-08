package repository

import "Trovero/Backend/internal/models"

type UserRepository interface {
	CreateUser(user *models.User)error
	GetUserByEmail(email string)(*models.User,error)
	GetAll()([]models.User,error)
	GetUserByID(UserID int64)(*models.User,error)
	DeleteUserByID(UserID int64)error
}

