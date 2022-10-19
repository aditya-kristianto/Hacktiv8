package postgres

import (
	"final-project/internal/app/model"
	"final-project/internal/app/user/repository"

	"github.com/google/uuid"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type (
	Repository struct {
		db *gorm.DB
	}
)

func NewRepository(db *gorm.DB) repository.Repository {
	return &Repository{
		db: db,
	}
}

func (u *Repository) Create(data *model.User) (*model.User, error) {
	err := u.db.Create(data).Error
	if err != nil {
		return nil, err
	}

	return data, nil
}

func (u *Repository) Read(email string) (*model.User, error) {
	var user model.User
	err := u.db.Model(&model.User{}).Where("email = ?", email).Find(&user).Error
	if err != nil {
		return nil, err
	}

	return &user, nil
}

func (u *Repository) ReadByID(userID *uuid.UUID) (*model.User, error) {
	var user model.User
	err := u.db.Model(&model.User{}).Where("id = ?", userID).Find(&user).Error
	if err != nil {
		return nil, err
	}

	return &user, nil
}

func (u *Repository) Update(userID *uuid.UUID, data *model.User) (*model.User, error) {
	var user model.User
	err := u.db.Model(&user).Clauses(clause.Returning{}).Where("id = ?", userID).Updates(data).Error
	if err != nil {
		return nil, err
	}

	return &user, nil
}

func (u *Repository) Delete(userID *uuid.UUID) (*model.User, error) {
	var user model.User
	err := u.db.Where("id = ?", userID).Delete(&user).Error
	if err != nil {
		return nil, err
	}

	return &user, nil
}
