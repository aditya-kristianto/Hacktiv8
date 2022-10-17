package gorm

import (
	"sesi7/model"
	"sesi7/repository"

	"gorm.io/gorm"
)

type userRepo struct {
	db *gorm.DB
}

func NewUserRepo(db *gorm.DB) repository.UserRepository {
	return &userRepo{
		db: db,
	}
}

func (u *userRepo) CreateUser(user *model.User) error {
	return u.db.Create(user).Error
}

func (u *userRepo) GetUsers() (*[]model.User, error) {
	var users []model.User

	err := u.db.Find(&users).Error
	if err != nil {
		return nil, err
	}

	return &users, nil
}
