package model

import (
	"time"

	"github.com/google/uuid"
)

type (
	User struct {
		ID        uuid.UUID `json:"id" gorm:"primaryKey;type:uuid;default:gen_random_uuid()"`
		Username  string    `json:"username" validate:"required" gorm:"uniqueIndex"`
		Email     string    `json:"email" validate:"email,required" gorm:"uniqueIndex"`
		Password  string    `json:"password" validate:"required,gte=6"`
		Age       int       `json:"age" validate:"required,gt=8"`
		CreatedAt time.Time `json:"created_at" validate:"required"`
		UpdatedAt time.Time `json:"updated_at" validate:"required"`
	}
)
