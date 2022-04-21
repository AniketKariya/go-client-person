package models

import "gorm.io/gorm"

type Person struct {
	gorm.Model
	Name    string
	Hobbies string
}
