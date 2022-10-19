package comment

import (
	"time"

	"github.com/google/uuid"
)

type (
	CreateResponse struct {
		ID        uuid.UUID `json:"id"`
		Message   string    `json:"message"`
		PhotoID   uuid.UUID `json:"photo_id"`
		UserID    uuid.UUID `json:"user_id"`
		CreatedAt time.Time `json:"created_at"`
	}

	GetResponse struct {
		ID        uuid.UUID    `json:"id"`
		UserID    uuid.UUID    `json:"user_id"`
		PhotoID   uuid.UUID    `json:"photo_id"`
		Message   string       `json:"message"`
		CreatedAt time.Time    `json:"created_at"`
		UpdatedAt time.Time    `json:"updated_at"`
		User      CommentUser  `json:"user"`
		Photo     CommentPhoto `json:photo`
	}

	CommentUser struct {
		ID       uuid.UUID `json:"id"`
		Username string    `json:"username"`
		Email    string    `json:"email"`
	}

	CommentPhoto struct {
		ID       uuid.UUID `json:"id"`
		Title    string    `json:"title"`
		Caption  string    `json:"caption"`
		PhotoURL string    `json:"photo_url"`
		UserID   uuid.UUID `json:"user_id"`
	}

	UpdateResponse struct {
		ID        uuid.UUID `json:"id"`
		UserID    uuid.UUID `json:"user_id"`
		PhotoID   uuid.UUID `json:"photo_id"`
		Message   string    `json:"message"`
		UpdatedAt time.Time `json:"updated_at"`
	}
)
