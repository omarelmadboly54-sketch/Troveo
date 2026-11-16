package models


type Category struct{
	ID				int64				`gorm:"primarykey" json:"id"`
	Category		string				`gorm:"not null" json:"category"`
	ParentID		*int64				`json:"parent_id"`
}

