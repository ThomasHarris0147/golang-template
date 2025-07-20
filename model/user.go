package model

import "gorm.io/gorm"

type Role string

const (
	Admin       Role = "admin"
	GeneralUser Role = "user"
)

type User struct {
	gorm.Model
	Name     string `json:"name"`
	Email    string `json:"email" gorm:"unique"`
	Password string `json:"password"`
	Role     Role   `json:"role" gorm:"default: 'user'"`
}
