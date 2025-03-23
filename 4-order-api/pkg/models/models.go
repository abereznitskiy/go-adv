package models

import (
	"github.com/lib/pq"
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	PhoneNumber string
	SessionId   string
	Code        string
	Orders      []*Order `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
}

type Order struct {
	gorm.Model
	UserID   uint       `json:"user_id"`
	Products []*Product `gorm:"many2many:order_products;constraint:OnDelete:CASCADE;"`
}

type Product struct {
	gorm.Model
	Name        string
	Description string
	Images      pq.StringArray `gorm:"type:text[]"`
}
