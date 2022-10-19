package repository

import (
	"final-project/internal/app/model"

	"github.com/google/uuid"
)

type (
	Repository interface {
		Create(*model.Comment) (*model.Comment, error)
		Read(*uuid.UUID) (*[]model.Comment, error)
		ReadByID(*uuid.UUID, *uuid.UUID) (*model.Comment, error)
		Update(*uuid.UUID, *model.Comment) (*model.Comment, error)
		Delete(*uuid.UUID) (*model.Comment, error)
	}
)
