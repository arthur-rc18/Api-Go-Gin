package models

import "gorm.io/gorm"

type Knight struct {
	gorm.Model
	Name            string `json:"name"`
	Weapon          string `json:"weapon"`
	Characteristics string `json:"caracteristic"`
}
