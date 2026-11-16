package repository

import "Trovero/Backend/internal/models"

type ProductRepository interface {
	CreateProduct(product *models.Product)error
	GetAll(page,limit int)([]models.Product,error)
	GetProductByID(id int64)(*models.Product,error)
	GetMyProducts(UserID int64,page,limit int)([]models.Product,error)
	GetMyProduct(id,userID int64)(*models.Product,error)
	DeleteProductByID(id,UserID int64)error
	GetProductsByCategory(CategoryID int64)([]models.Product,error)
}


