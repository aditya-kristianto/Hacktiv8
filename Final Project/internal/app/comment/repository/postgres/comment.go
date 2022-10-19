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

	return data, nil
}

func (c *Repository) Read(userID *uuid.UUID) (*[]model.Comment, error) {
	var comments []model.Comment
	err := c.db.Raw("SELECT comments.*, \"User\".id AS \"User__id\", \"User\".email AS \"User__email\", \"User\".username AS \"User__username\", \"Photo\".id AS \"Photo__id\", \"Photo\".title AS \"Photo__title\", \"Photo\".caption AS \"Photo__caption\", \"Photo\".photo_url AS \"Photo__photo_url\", \"Photo\".user_id AS \"Photo__user_id\" FROM comments LEFT JOIN users \"User\" ON comments.user_id = \"User\".id LEFT JOIN photos \"Photo\" ON comments.photo_id = \"Photo\".id WHERE comments.user_id = ?", userID).Scan(&comments).Error

	// err := c.db.Model(&model.Comment{}).Where("user_id = ?", userID).Find(&comments).Error
	if err != nil {
		return nil, err
	}

	return &comments, nil
}

func (c *Repository) ReadByID(commentID *uuid.UUID, userID *uuid.UUID) (*model.Comment, error) {
	var comment model.Comment
	err := c.db.Model(&model.Comment{}).Where("id = ? and user_id = ?", commentID, userID).Find(&comment).Error
	if err != nil {
		return nil, err
	}

	return &comment, nil
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
