package gorm

import (
	"sesi7/gin/repository"

	"gorm.io/gorm"
)

type userGorm struct {
	db *gorm.DB
}

func NewUserRepo(db *gorm.DB) repository.UserRepo {
	return &userGorm{
		db: db,
	}
}

func (u *userGorm) CreateUser(user *repository.User) error {

	return u.db.Create(user).Error
}
func (u *userGorm) GetUsers() (*[]repository.User, error) {

	var users []repository.User
	err := u.db.Find(&users).Error
	if err != nil {
		return nil, err
	}
	return &users, nil
}
func (u *userGorm) GetUserByID(id int) (*repository.User, error) {
	return nil, nil
}
