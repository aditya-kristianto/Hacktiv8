package repository

import (
	"final-project/internal/app/model"

	"github.com/google/uuid"
)

type (
	Repository interface {
		Create(*model.User) (*model.User, error)
		Read(string) (*model.User, error)
		ReadByID(userID *uuid.UUID) (*model.User, error)
		Update(*uuid.UUID, *model.User) (*model.User, error)
		Delete(*uuid.UUID) (*model.User, error)
	}
)
