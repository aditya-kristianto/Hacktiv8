package postgres

import (
	"final-project/internal/app/model"
	"final-project/internal/app/photo/repository"

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

func (p *Repository) Create(data *model.Photo) (*model.Photo, error) {
	err := p.db.Create(data).Error
	if err != nil {
		return nil, err
	}

	return data, nil
}

func (p *Repository) Read(userID *uuid.UUID) (*[]model.Photo, error) {
	var photos []model.Photo
	err := p.db.Model(&model.Photo{}).Where("user_id = ?", userID).Find(&photos).Error
	if err != nil {
		return nil, err
	}

	return &photos, nil
}

func (p *Repository) Update(photoID *uuid.UUID, data *model.Photo) (*model.Photo, error) {
	var photo model.Photo
	err := p.db.Model(&photo).Clauses(clause.Returning{}).Where("id = ?", photoID).Updates(data).Error
	if err != nil {
		return nil, err
	}

	return &photo, nil
}

func (p *Repository) Delete(photoID *uuid.UUID) (*model.Photo, error) {
	var photo model.Photo
	err := p.db.Where("id = ?", photoID).Delete(&photo).Error
	if err != nil {
		return nil, err
	}

	return &photo, nil
}
