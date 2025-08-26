package models

import "gorm.io/gorm"

type User struct {
	gorm.Model
	UserID   uint   `json:"id"`
	Name     string `json:"name"`
	Password string `json:"password"`
}
