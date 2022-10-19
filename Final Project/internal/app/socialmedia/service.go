package socialmedia

import (
	"errors"
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

func (s *Service) GetSocialmedia(userID *uuid.UUID) (*[]GetResponse, error) {
	socialmedias, err := s.repository.Read(userID)
	if err != nil {
		return nil, err
	}

	var resp []GetResponse
	for _, item := range *socialmedias {
		socialmedia := GetResponse{
			ID:             item.ID,
			Name:           item.Name,
			SocialMediaURL: item.SocialMediaURL,
			UserID:         item.UserID,
			CreatedAt:      item.CreatedAt,
			UpdatedAt:      item.UpdatedAt,
			User: SocialmediaUser{
				ID:       item.User.ID,
				Username: item.User.Username,
			},
		}
		resp = append(resp, socialmedia)
	}
	return &resp, nil
}

func (s *Service) UpdateSocialmedia(socialmediaID *uuid.UUID, userID *uuid.UUID, req *SocialMediaRequest) (*UpdateResponse, error) {
	isExistSocialmedia, err := s.repository.ReadByID(socialmediaID, userID)
	if err != nil {
		return nil, err
	}

	if isExistSocialmedia.ID == uuid.Nil {
		return nil, errors.New("social media id is not registered")
	}

	data := model.SocialMedia{
		Name:           req.Name,
		SocialMediaURL: req.SocialMediaURL,
	}
	socialmedia, err := s.repository.Update(socialmediaID, userID, &data)
	if err != nil {
		return nil, err
	}

	resp := &UpdateResponse{
		ID:             socialmedia.ID,
		Name:           socialmedia.Name,
		SocialMediaURL: socialmedia.SocialMediaURL,
		UserID:         socialmedia.UserID,
		UpdatedAt:      socialmedia.UpdatedAt,
	}
	return resp, nil
}

func (s *Service) DeleteSocialmedia(socialmediaID *uuid.UUID, userID *uuid.UUID) error {
	isExistSocialmedia, err := s.repository.ReadByID(socialmediaID, userID)
	if err != nil {
		return err
	}

	if isExistSocialmedia.ID == uuid.Nil {
		return errors.New("social media id is not registered")
	}

	_, err = s.repository.Delete(socialmediaID)
	if err != nil {
		return err
	}

	return nil
}
