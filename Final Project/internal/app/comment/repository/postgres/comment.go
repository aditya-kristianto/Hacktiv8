package postgres

import (
	"final-project/internal/app/comment/repository"
	"final-project/internal/app/model"

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

func (c *Repository) Create(data *model.Comment) (*model.Comment, error) {
	err := c.db.Create(data).Error
	if err != nil {
		return nil, err
	}

	return nil, nil
}

func (c *Repository) Read(userID *uuid.UUID) (*[]model.Comment, error) {
	var comments []model.Comment
	err := c.db.Model(&model.Comment{}).Where("user_id = ?", userID).Find(&comments).Error
	if err != nil {
		return nil, err
	}

	return &comments, nil
}

func (c *Repository) Update(commentID *uuid.UUID, data *model.Comment) (*model.Comment, error) {
	var comment model.Comment
	err := c.db.Model(&comment).Clauses(clause.Returning{}).Where("id = ?", commentID).Updates(data).Error
	if err != nil {
		return nil, err
	}

	return &comment, nil
}

func (c *Repository) Delete(commentID *uuid.UUID) (*model.Comment, error) {
	var comment model.Comment
	err := c.db.Where("id = ?", commentID).Delete(&comment).Error
	if err != nil {
		return nil, err
	}

	return &comment, nil
}
