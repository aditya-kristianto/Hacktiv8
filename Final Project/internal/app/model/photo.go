package model

import (
	"time"

	"github.com/google/uuid"
)

type (
	Photo struct {
		ID        uuid.UUID `gorm:"primaryKey;type:uuid;default:gen_random_uuid()"`
		Title     string    `json:"title" validate:"required"`
		Caption   string    `json:"caption" validate:"required"`
		PhotoURL  string    `json:"photo_url" validate:"required"`
		CreatedAt time.Time
		UpdatedAt time.Time
	}
)
