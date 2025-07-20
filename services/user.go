package services

import "gorm.io/gorm"

type IUserService interface {
}

type UserService struct {
	db *gorm.DB
}

func NewUserService(db *gorm.DB) IUserService {
	return &UserService{db: db}
}
