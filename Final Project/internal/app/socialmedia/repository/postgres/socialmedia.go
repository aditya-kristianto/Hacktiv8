package postgres

import (
	"final-project/internal/app/model"
	"final-project/internal/app/socialmedia/repository"

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

func (s *Repository) Create(data *model.SocialMedia) (*model.SocialMedia, error) {
	err := s.db.Create(data).Error
	if err != nil {
		return nil, err
	}

	return data, nil
}

func (s *Repository) Read(userID *uuid.UUID) (*[]model.SocialMedia, error) {
	var socialmedias []model.SocialMedia
	err := s.db.Raw("SELECT social_media.*, \"User\".id AS \"User__id\", \"User\".username AS \"User__username\" FROM social_media LEFT JOIN users \"User\" ON social_media.user_id = \"User\".id WHERE user_id = ?", userID).Scan(&socialmedias).Error
	if err != nil {
		return nil, err
	}

	return &socialmedias, nil
}

func (s *Repository) ReadByID(socialmediaID *uuid.UUID, userID *uuid.UUID) (*model.SocialMedia, error) {
	var socialmedia model.SocialMedia
	err := s.db.Model(&socialmedia).Where("id = ? and user_id = ?", socialmediaID, userID).Scan(&socialmedia).Error
	if err != nil {
		return nil, err
	}

	return &socialmedia, nil
}

func (s *Repository) Update(socialmediaID *uuid.UUID, userID *uuid.UUID, data *model.SocialMedia) (*model.SocialMedia, error) {
	var socialmedia model.SocialMedia
	err := s.db.Model(&socialmedia).Clauses(clause.Returning{}).Where("id = ? and user_id = ?", socialmediaID, userID).Updates(data).Error
	if err != nil {
		return nil, err
	}

	return &socialmedia, nil
}

func (s *Repository) Delete(socialmediaID *uuid.UUID) (*model.SocialMedia, error) {
	var socialmedia model.SocialMedia
	err := s.db.Where("id = ?", socialmediaID).Delete(&socialmedia).Error
	if err != nil {
		return nil, err
	}

	return &socialmedia, nil
}
