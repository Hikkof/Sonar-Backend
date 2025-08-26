package models

import "gorm.io/gorm"

type Cart struct {
	gorm.Model
	CartID   uint      `json:"id"`
	Contents []Product `json:"contents" gorm:"foreignKey:ProductID"`
}
