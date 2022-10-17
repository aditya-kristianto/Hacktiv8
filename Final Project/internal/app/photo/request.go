package photo

type (
	PhotoRequest struct {
		Title    string `json:"title" example:"echo_golang"`
		Caption  string `json:"caption" example:"echo"`
		PhotoURL string `json:"photo_url" example:"https://cdn.labstack.com/images/echo-logo.svg"`
	}
)
