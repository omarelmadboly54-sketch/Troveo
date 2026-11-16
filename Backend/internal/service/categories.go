package service

import (
	"Trovero/Backend/internal/apperrors"
	"Trovero/Backend/internal/models"
	"Trovero/Backend/internal/repository"
	"errors"

	"gorm.io/gorm"
)

type categoryService struct {
	categoryRepository repository.CategoryRepository
}

type CategoryService interface {
	CreateCategory(category *models.Category) error
	GetAll()([]models.Category,error)
	DeleteCategoryByID(id int64)error
}

func NewCategoryService(categoryRepository repository.CategoryRepository) CategoryService {
	return &categoryService{
		categoryRepository: categoryRepository,
	}
}

func(s*categoryService)CreateCategory(category *models.Category) error{
	return s.categoryRepository.CreateCategory(category)
}

func(s*categoryService)GetAll()([]models.Category,error){
	return s.categoryRepository.GetAll()
}

func(s *categoryService)DeleteCategoryByID(id int64)error{
	err:=s.categoryRepository.DeleteCategoryByID(id)
	if err!=nil{
		if errors.Is(err,gorm.ErrRecordNotFound){
			return apperrors.ErrCategoryNotFound
		}
		return err
	}

	return nil
}
