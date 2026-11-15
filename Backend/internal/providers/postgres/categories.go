package postgres

import (
	"Trovero/Backend/internal/models"
	"Trovero/Backend/internal/repository"

	"gorm.io/gorm"
)

type categoryRepository struct {
	db *gorm.DB
}

func NewCategoryRepository(db *gorm.DB) repository.CategoryRepository {
	return &categoryRepository{
		db: db,
	}
}

func (r *categoryRepository) CreateCategory(category *models.Category) error{
	return r.db.Create(category).Error
}


func (r *categoryRepository)GetAll()([]models.Category,error){
	var category []models.Category

	result:=r.db.Find(&category)
	if result.Error!=nil{
		return []models.Category{},result.Error
	}

	return category,nil
}


func(r *categoryRepository)DeleteCategoryByID(id int64)error{
	

	result:=r.db.Where("id=?",id).Delete(&models.Category{})
	if result.Error!=nil{
		return result.Error
	}

	if result.RowsAffected==0{
		return gorm.ErrRecordNotFound
	}

	return nil

}
