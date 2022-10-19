package model

import (
	"time"

	"github.com/google/uuid"
)

type (
	Comment struct {
		ID        uuid.UUID `json:"id" gorm:"primaryKey;type:uuid;default:gen_random_uuid()"`
		UserID    uuid.UUID `json:"user_id"`
		PhotoID   uuid.UUID `json:"photo_id"`
		Message   string    `json:"message" validate:"required"`
		CreatedAt time.Time `json:"created_at"`
		UpdatedAt time.Time `json:"updated_at"`
		User      User      `json:"-" gorm:"foreignKey:UserID;references:ID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
		Photo     Photo     `json:"-" gorm:"foreignKey:PhotoID;references:ID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	}
)
