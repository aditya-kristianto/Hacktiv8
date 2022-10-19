package photo

import (
	"gorm.io/gorm"

	"final-project/internal/app/model"
	"final-project/internal/app/photo/repository"
	"final-project/internal/app/photo/repository/postgres"

	"github.com/google/uuid"
)

type (
	Service struct {
		repository repository.Repository
	}
)

func NewService(db *gorm.DB) *Service {
	return &Service{
		repository: postgres.NewRepository(db),
	}
}

func (p *Service) CreatePhoto(userID *uuid.UUID, req *PhotoRequest) (*model.Photo, error) {
	data := &model.Photo{
		Title:    req.Title,
		Caption:  req.Caption,
		PhotoURL: req.PhotoURL,
		UserID:   *userID,
	}
	photo, err := p.repository.Create(data)
	if err != nil {
		return nil, err
	}

	return photo, nil
}

func (p *Service) GetPhoto(userID *uuid.UUID) (*[]model.Photo, error) {
	photos, err := p.repository.Read(userID)
	if err != nil {
		return nil, err
	}

	return photos, nil
}

func (p *Service) UpdatePhoto(photoID *uuid.UUID, data *model.Photo) (*model.Photo, error) {
	photo, err := p.repository.Update(photoID, data)
	if err != nil {
		return nil, err
	}

	return photo, nil
}

func (p *Service) DeletePhoto(photoID *uuid.UUID) error {
	_, err := p.repository.Delete(photoID)
	if err != nil {
		return err
	}

	return nil
}
