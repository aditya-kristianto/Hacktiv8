package model

import (
	"time"

	"github.com/google/uuid"
)

type (
	SocialMedia struct {
		ID             uuid.UUID `json:"id" gorm:"primaryKey;type:uuid;default:gen_random_uuid()"`
		Name           string    `json:"name" validate:"required"`
		SocialMediaURL string    `json:"social_media_url" validate:"required"`
		UserID         uuid.UUID `json:"user_id"`
		CreatedAt      time.Time `json:"created_at"`
		UpdatedAt      time.Time `json:"updated_at"`
		User           User      `gorm:"foreignKey:UserID;references:ID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	}
)
