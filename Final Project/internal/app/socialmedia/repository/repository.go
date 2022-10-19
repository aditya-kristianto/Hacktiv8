package repository

import (
	"final-project/internal/app/model"

	"github.com/google/uuid"
)

type (
	Repository interface {
		Create(*model.SocialMedia) (*model.SocialMedia, error)
		Read(*uuid.UUID) (*[]model.SocialMedia, error)
		ReadByID(*uuid.UUID, *uuid.UUID) (*model.SocialMedia, error)
		Update(*uuid.UUID, *uuid.UUID, *model.SocialMedia) (*model.SocialMedia, error)
		Delete(*uuid.UUID) (*model.SocialMedia, error)
	}
)
