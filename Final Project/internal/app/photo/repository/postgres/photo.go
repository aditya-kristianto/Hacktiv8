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
	err := p.db.Raw("SELECT photos.*, NULL AS \"User__id\", \"User\".email AS \"User__email\", \"User\".username AS \"User__username\" FROM photos LEFT JOIN users \"User\" ON photos.user_id = \"User\".id WHERE user_id = ?", userID).Scan(&photos).Error
	if err != nil {
		return nil, err
	}

	return &photos, nil
}

func (p *Repository) ReadByPhotoID(photoID *uuid.UUID, userID *uuid.UUID) (*model.Photo, error) {
	var photo model.Photo
	err := p.db.Model(&photo).Where("id = ? and user_id = ?", photoID, userID).Scan(&photo).Error
	if err != nil {
		return nil, err
	}

	return &photo, nil
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
