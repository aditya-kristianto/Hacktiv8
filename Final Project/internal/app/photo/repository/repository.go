package repository

import (
	"final-project/internal/app/model"

	"github.com/google/uuid"
)

type (
	Repository interface {
		Create(*model.Photo) (*model.Photo, error)
		Read(*uuid.UUID) (*[]model.Photo, error)
		ReadByPhotoID(*uuid.UUID, *uuid.UUID) (*model.Photo, error)
		Update(*uuid.UUID, *model.Photo) (*model.Photo, error)
		Delete(*uuid.UUID) (*model.Photo, error)
	}
)
