package models

import "time"

type User struct {
	ID          int64  		`gorm:"primarykey" json:"id"`
	FirstName   string 		`gorm:"not null" json:"first_name"`
	LastName    string 		`gorm:"not null" json:"last_name"`
	Email       string 		`gorm:"not null;unique" json:"email"`
	Password    string 		`gorm:"not null" json:"-"`
	Role        string 		`gorm:"not null;default:'user'" json:"role"`
	AccountType string 		`gorm:"not null" json:"account_type"`
	CreatedAt   time.Time	`gorm:"autoCreateTime" json:"-"`
	UpdatedAt	time.Time	`gorm:"autoUpdateTime" json:"-"`
}

