package models


type Review struct{
	ID			int64		`gorm:"primarykey" json:"id"`
	ProductID	int64		`gorm:"not null" json:"product_id"`
	
}