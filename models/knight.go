package models

import (
	"gopkg.in/validator.v2"
	"gorm.io/gorm"
)

type Knight struct {
	gorm.Model
	// Specifying the field's characteristics using validate
	Name            string `json:"name" validate:"nonzero"`
	Weapon          string `json:"weapon" validate:"len=20"`
	Characteristics string `json:"characteristic" validate:"len=1000"`
	Age             string `json:"age" validate:"len=4, regexp=^[0-9]*$" swaggertype:"string"`
}

func ValidateKnight(knight *Knight) error {
	if err := validator.Validate(knight); err != nil {
		return err
	}
	return nil
}
