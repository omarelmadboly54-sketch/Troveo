package presenters

import "Trovero/Backend/internal/models"

type CreateUserReq struct {
	FirstName   string `json:"first_name" binding:"required"`
	LastName    string `json:"last_name" binding:"required"`
	Email       string `json:"email" binding:"required,email"`
	Password    string `json:"password" binding:"required,min=8"`
	AccountType string `json:"account_type" binding:"required,oneof=seller customer"`
}

type LoginReq struct{
	Email			string		`json:"email" binding:"required"`
	Password		string		`json:"password" binding:"required"`
}


func (c *CreateUserReq) ToUser() models.User{
	return models.User{
		FirstName: c.FirstName,
		LastName: c.LastName,
		Email: c.Email,
		Password: c.Password,
		AccountType: c.AccountType,
	}
}
