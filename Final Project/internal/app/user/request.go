package user

type (
	RegisterRequest struct {
		Age      int    `json:"age" example:"31"`
		Email    string `json:"email" example:"aditya.kristianto@mncgroup.com"`
		Password string `json:"password" example:"hacktiv8"`
		Username string `json:"username" example:"aditya.kristianto"`
	}

	LoginRequest struct {
		Email    string `json:"email" example:"aditya.kristianto@mncgroup.com"`
		Password string `json:"password" example:"hacktiv8"`
	}

	UpdateUserRequest struct {
		Email    string `json:"email" example:"aditya.kristianto@mncgroup.com"`
		Username string `json:"username" example:"aditya.kristianto"`
	}
)
