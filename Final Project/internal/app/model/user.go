package model

import (
	"time"

	"github.com/google/uuid"
)

type (
	User struct {
		ID        uuid.UUID `json:"id,omitempty" gorm:"primaryKey;type:uuid;default:gen_random_uuid()"`
		Username  string    `json:"username" validate:"required" gorm:"uniqueIndex"`
		Email     string    `json:"email" validate:"email,required" gorm:"uniqueIndex"`
		Password  string    `json:"-" validate:"required,gte=6"`
		Age       int       `json:"age,omitempty" validate:"required,gt=8"`
		CreatedAt time.Time `json:"-" validate:"required"`
		UpdatedAt time.Time `json:"updated_at,omitempty" validate:"required"`
	}
)
