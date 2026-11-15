package repository

import "Trovero/Backend/internal/models"

type CategoryRepository interface {
	CreateCategory(category *models.Category) error
	GetAll()([]models.Category,error)
	DeleteCategoryByID(id int64)error
}


