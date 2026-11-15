package presenters

import "Trovero/Backend/internal/models"

type CreateProductRequest struct {
	Name        string  `form:"name" binding:"required"`
	Description string  `form:"description" binding:"required"`
	Price       float64 `form:"price" binding:"required,gt=0"`
	Quantity    int     `form:"quantity" binding:"gte=0"`
	CategoryID  int64   `form:"category_id" binding:"required,gt=0"`
}

func (c *CreateProductRequest) ToProduct() models.Product{
	return models.Product{
		Name: c.Name,
		Description: c.Description,
		Price: c.Price,
		Quantity: c.Quantity,
		CategoryID: c.CategoryID,
	}
}
