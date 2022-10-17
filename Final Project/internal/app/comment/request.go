package comment

type (
	CreateCommentRequest struct {
		PhotoID string `json:"photo_id"`
		Message string `json:"message" example:"required"`
	}

	UpdateCommentRequest struct {
		Message string `json:"message" example:"required"`
	}
)
