package models

import "gorm.io/gorm"

type M5 struct { // nazwa do zmiany
	gorm.Model
	ID   uint `json:"id"`
	Name int  `json:"name"`
}
