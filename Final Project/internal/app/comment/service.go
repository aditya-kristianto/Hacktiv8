package comment

import (
	"errors"
	"final-project/internal/app/comment/repository"
	"final-project/internal/app/comment/repository/postgres"
	"final-project/internal/app/model"
	photoRepository "final-project/internal/app/photo/repository"
	photoPostgresRepository "final-project/internal/app/photo/repository/postgres"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type (
	Service struct {
		repository      repository.Repository
		photoRepository photoRepository.Repository
	}
)

func NewService(db *gorm.DB) *Service {
	return &Service{
		repository:      postgres.NewRepository(db),
		photoRepository: photoPostgresRepository.NewRepository(db),
	}
}

func (c *Service) CreateComment(userID *uuid.UUID, req *CreateCommentRequest) (*CreateResponse, error) {
	photoID, err := uuid.Parse(req.PhotoID)
	if err != nil {
		return nil, err
	}

	photo, err := c.photoRepository.ReadByPhotoID(&photoID, userID)
	if err != nil {
		return nil, err
	}

	if photo.ID == uuid.Nil {
		return nil, errors.New("photo id is not registered")
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

	resp := &CreateResponse{
		ID:        comment.ID,
		Message:   comment.Message,
		PhotoID:   comment.PhotoID,
		UserID:    comment.UserID,
		CreatedAt: comment.CreatedAt,
	}
	return resp, nil
}

func (c *Service) GetComment(userID *uuid.UUID) (*[]GetResponse, error) {
	comments, err := c.repository.Read(userID)
	if err != nil {
		return nil, err
	}

	var resp []GetResponse
	for _, item := range *comments {
		comment := GetResponse{
			ID:        item.ID,
			UserID:    item.UserID,
			PhotoID:   item.PhotoID,
			Message:   item.Message,
			CreatedAt: item.CreatedAt,
			UpdatedAt: item.UpdatedAt,
			User: CommentUser{
				ID:       item.User.ID,
				Username: item.User.Username,
				Email:    item.User.Email,
			},
			Photo: CommentPhoto{
				ID:       item.Photo.ID,
				Title:    item.Photo.Title,
				Caption:  item.Photo.Caption,
				PhotoURL: item.Photo.PhotoURL,
				UserID:   item.Photo.UserID,
			},
		}
		resp = append(resp, comment)
	}
	return &resp, nil
}

func (c *Service) UpdateComment(commentID *uuid.UUID, data *model.Comment) (*UpdateResponse, error) {
	comment, err := c.repository.Update(commentID, data)
	if err != nil {
		return nil, err
	}

	resp := &UpdateResponse{
		ID:        comment.ID,
		UserID:    comment.UserID,
		PhotoID:   comment.PhotoID,
		Message:   comment.Message,
		UpdatedAt: comment.UpdatedAt,
	}
	return resp, nil
}

func (c *Service) DeleteComment(commentID *uuid.UUID, userID *uuid.UUID) error {
	comment, err := c.repository.ReadByID(commentID, userID)
	if err != nil {
		return err
	}

	if comment.ID == uuid.Nil {
		return errors.New("comment id is not registered")
	}

	_, err = c.repository.Delete(commentID)
	if err != nil {
		return err
	}

	return nil
}
