package models

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Id	int64 `json:"id" gorm:"primaryKey"`
	Name string `json:"name"`
	Email string `json:"email"`
	Password string `json:"password"`
}