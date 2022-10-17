package socialmedia

type (
	SocialMediaRequest struct {
		Name           string `json:"name" example:"linkedin"`
		SocialMediaURL string `json:"social_media_url" example:"https://www.linkedin.com/in/aditya-kristianto"`
	}
)
