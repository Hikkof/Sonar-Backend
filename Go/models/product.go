package models

import (
	"gorm.io/gorm"
)

type Product struct {
	gorm.Model
	ProductID   uint    `json:"id"`
	Name        string  `json:"name"`
	Price       float32 `json:"price"`
	Description string  `json:"description"`
	SellerID    uint    `json:"seller_id"`
}
