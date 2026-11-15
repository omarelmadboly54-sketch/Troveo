package repository

import "Trovero/Backend/internal/models"

type ProductRepository interface {
	CreateProduct(product *models.Product)error
}