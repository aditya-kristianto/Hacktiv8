package comment

import (
	"final-project/internal/app/comment/repository"
	"final-project/internal/app/comment/repository/postgres"
	"final-project/internal/app/model"

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

func (c *Service) CreateComment(userID *uuid.UUID, req *CreateCommentRequest) (*model.Comment, error) {
	photoID, err := uuid.Parse(req.PhotoID)
	if err != nil {
		return nil, err
	}

	data := &model.Comment{
		Message: req.Message,
		PhotoID: photoID,
		UserID:  *userID,
	}
	comment, err := c.repository.Create(data)
	if err != nil {
		return nil, err
	}

	return comment, nil
}

func (c *Service) GetComment(userID *uuid.UUID) (*[]model.Comment, error) {
	comments, err := c.repository.Read(userID)
	if err != nil {
		return nil, err
	}

	return comments, nil
}

func (c *Service) UpdateComment(commentID *uuid.UUID, data *model.Comment) (*model.Comment, error) {
	comment, err := c.repository.Update(commentID, data)
	if err != nil {
		return nil, err
	}

	return comment, nil
}

func (c *Service) DeleteComment(commentID *uuid.UUID) error {
	_, err := c.repository.Delete(commentID)
	if err != nil {
		return err
	}

	return nil
}
