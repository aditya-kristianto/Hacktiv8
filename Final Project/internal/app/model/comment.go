package model

import (
	"time"

	"github.com/google/uuid"
)

type (
	Comment struct {
		ID        uuid.UUID `gorm:"primaryKey;type:uuid;default:gen_random_uuid()"`
		UserID    uuid.UUID `json:"user_id"`
		PhotoID   uuid.UUID `json:"photo_id"`
		Message   string    `json:"message" validate:"required"`
		CreatedAt time.Time
		UpdatedAt time.Time
		User      User  `gorm:"foreignKey:UserID;references:ID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
		Photo     Photo `gorm:"foreignKey:PhotoID;references:ID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	}
)
