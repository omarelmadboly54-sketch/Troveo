package presenters

import "Trovero/Backend/internal/models"

type CreateAddressReq struct {
	City            string `json:"city" binding:"required"`
	Area            string `json:"area" binding:"required"`
	Street          string `json:"street" binding:"required"`
	BuildingNumber  int    `json:"building_number" binding:"required"`
	ApartmentNumber *int    `json:"apartment_number"`
	PhoneNumber     string `json:"phone_number" binding:"required"`
}

func (c *CreateAddressReq) ToAddress() models.Address{
	return models.Address{
		City: c.City,
		Area: c.Area,
		Street: c.Street,
		BuildingNumber: c.BuildingNumber,
		ApartmentNumber: c.ApartmentNumber,
		PhoneNumber: c.PhoneNumber,
	}
}


