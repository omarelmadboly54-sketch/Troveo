package postgres

import (
	"Trovero/Backend/internal/models"
	"Trovero/Backend/internal/repository"

	"gorm.io/gorm"
)

type userRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) repository.UserRepository {
	return &userRepository{
		db: db,
	}
}

func (r *userRepository) CreateUser(user *models.User) error {
	return r.db.Create(&user).Error
}

func (r *userRepository) GetUserByEmail(email string) (*models.User, error) {
	var user models.User

	result := r.db.Where("email=?", email).First(&user)
	if result.Error != nil {
		return nil, result.Error
	}
	return &user, nil
}

func (r *userRepository) GetUserByID(UserID int64) (*models.User, error) {
	var user models.User

	result := r.db.Where("id=?", UserID).First(&user)

	if result.Error != nil {
		return nil, result.Error
	}
	return &user, nil
}


func (r *userRepository)DeleteUserByID(UserID int64)error{
	result:= r.db.Where("id=? AND role=?",UserID,"user").Delete(&models.User{})

	if result.Error!=nil{
		return result.Error
	}

	if result.RowsAffected==0{
		return gorm.ErrRecordNotFound
	}
	return nil
}

func (r *userRepository)GetAll()([]models.User,error){
	var users []models.User

	result:=r.db.Where("id > ?",1).Find(&users)
	if result.Error!=nil{
		return []models.User{},result.Error
	}
	return users,nil
}

