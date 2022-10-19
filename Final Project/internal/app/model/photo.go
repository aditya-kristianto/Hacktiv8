package model

import (
	"time"

	"github.com/google/uuid"
)

type (
	Photo struct {
		ID        uuid.UUID `json:"id" gorm:"primaryKey;type:uuid;default:gen_random_uuid()"`
		Title     string    `json:"title" validate:"required"`
		Caption   string    `json:"caption" validate:"required"`
		PhotoURL  string    `json:"photo_url" validate:"required"`
		UserID    uuid.UUID `json:"user_id"`
		CreatedAt time.Time `json:"created_at"`
		UpdatedAt time.Time `json:"updated_at"`
		User      User      `json:"user" gorm:"foreignKey:UserID;references:ID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	}
)
