package socialmedia

import (
	"time"

	"github.com/google/uuid"
)

type (
	CreateResponse struct {
		ID             uuid.UUID `json:"id"`
		Name           string    `json:"name"`
		SocialMediaURL string    `json:"social_media_url"`
		UserID         uuid.UUID `json:"user_id"`
		CreatedAt      time.Time `json:"created_at"`
	}

	UpdateResponse struct {
		ID             uuid.UUID `json:"id"`
		Name           string    `json:"name"`
		SocialMediaURL string    `json:"social_media_url"`
		UserID         uuid.UUID `json:"user_id"`
		UpdatedAt      time.Time `json:"updated_at"`
	}

	GetResponse struct {
		ID             uuid.UUID       `json:"id"`
		Name           string          `json:"name"`
		SocialMediaURL string          `json:"social_media_url"`
		UserID         uuid.UUID       `json:"user_id"`
		CreatedAt      time.Time       `json:"created_at"`
		UpdatedAt      time.Time       `json:"updated_at"`
		User           SocialmediaUser `json:"user"`
	}

	SocialmediaUser struct {
		ID       uuid.UUID `json:"id"`
		Username string    `json:"username"`
	}
)
