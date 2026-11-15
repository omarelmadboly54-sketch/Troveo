package presenters

import "Trovero/Backend/internal/models"

type CreateCategoryReq struct {
	Category string `json:"category" binding:"required"`
	ParentID *int64 `json:"parent_id"`
}

func (c *CreateCategoryReq) ToCategory() models.Category{
	return models.Category{
		Category: c.Category,
		ParentID: c.ParentID,
	}
}
