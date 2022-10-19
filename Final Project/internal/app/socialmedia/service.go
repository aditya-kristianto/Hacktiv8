package socialmedia

import (
	"final-project/internal/app/model"
	"final-project/internal/app/socialmedia/repository"
	"final-project/internal/app/socialmedia/repository/postgres"

	"github.com/google/uuid"
	"gorm.io/gorm"
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

func (s *Service) CreateSocialmedia(userID *uuid.UUID, req *SocialMediaRequest) (*model.SocialMedia, error) {
	socialmedia, err := s.repository.Create(&model.SocialMedia{
		Name:           req.Name,
		SocialMediaURL: req.SocialMediaURL,
		UserID:         *userID,
	})
	if err != nil {
		return nil, err
	}

	return socialmedia, nil
}

func (s *Service) GetSocialmedia(userID *uuid.UUID) (*[]model.SocialMedia, error) {
	socialmedias, err := s.repository.Read(userID)
	if err != nil {
		return nil, err
	}

	return socialmedias, nil
}

func (s *Service) UpdateSocialmedia(socialmediaID *uuid.UUID, userID *uuid.UUID, req *SocialMediaRequest) (*model.SocialMedia, error) {
	data := model.SocialMedia{
		Name:           req.Name,
		SocialMediaURL: req.SocialMediaURL,
	}
	socialmedia, err := s.repository.Update(socialmediaID, userID, &data)
	if err != nil {
		return nil, err
	}

	return socialmedia, nil
}

func (s *Service) DeleteSocialmedia(socialmediaID *uuid.UUID) error {
	_, err := s.repository.Delete(socialmediaID)
	if err != nil {
		return err
	}

	return nil
}
