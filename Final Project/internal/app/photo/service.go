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

func (p *Service) CreatePhoto(userID *uuid.UUID, req *PhotoRequest) (*CreateResponse, error) {
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

	resp := &CreateResponse{
		ID:        photo.ID,
		Title:     photo.Title,
		Caption:   photo.Caption,
		PhotoURL:  photo.PhotoURL,
		UserID:    photo.UserID,
		CreatedAt: photo.CreatedAt,
	}
	return resp, nil
}

func (p *Service) GetPhoto(userID *uuid.UUID) (*[]model.Photo, error) {
	photos, err := p.repository.Read(userID)
	if err != nil {
		return nil, err
	}

	return photos, nil
}

func (p *Service) UpdatePhoto(photoID *uuid.UUID, data *model.Photo) (*UpdateResponse, error) {
	photo, err := p.repository.Update(photoID, data)
	if err != nil {
		return nil, err
	}

	resp := &UpdateResponse{
		ID:        photo.ID,
		Title:     photo.Title,
		Caption:   photo.Caption,
		PhotoURL:  photo.PhotoURL,
		UserID:    photo.UserID,
		UpdatedAt: photo.UpdatedAt,
	}
	return resp, nil
}

func (p *Service) DeletePhoto(photoID *uuid.UUID) error {
	_, err := p.repository.Delete(photoID)
	if err != nil {
		return err
	}

	return nil
}
