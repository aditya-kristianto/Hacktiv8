package photo

import (
	"time"

	"github.com/google/uuid"
)

type (
	CreateResponse struct {
		ID        uuid.UUID `json:"id"`
		Title     string    `json:"title" validate:"required"`
		Caption   string    `json:"caption" validate:"required"`
		PhotoURL  string    `json:"photo_url" validate:"required"`
		UserID    uuid.UUID `json:"user_id"`
		CreatedAt time.Time `json:"created_at,omitempty"`
	}

	UpdateResponse struct {
		ID        uuid.UUID `json:"id"`
		Title     string    `json:"title" validate:"required"`
		Caption   string    `json:"caption" validate:"required"`
		PhotoURL  string    `json:"photo_url" validate:"required"`
		UserID    uuid.UUID `json:"user_id"`
		UpdatedAt time.Time `json:"updated_at,omitempty"`
	}
)
