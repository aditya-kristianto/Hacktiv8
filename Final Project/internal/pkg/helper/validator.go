package helper

import (
	"net/http"

	"github.com/go-playground/validator"
	"github.com/labstack/echo/v4"
)

type (
	CustomValidator struct {
		validator *validator.Validate
	}
)

func (cv *CustomValidator) Validate(i interface{}) error {
	if err := cv.validator.Struct(i); err != nil {
		// Optionally, you could return the error to give each route more control over the status code
		resp := new(Response)
		resp.Status = http.StatusBadRequest
		resp.Message = "Bad Request"
		resp.Error = err.Error()
		return echo.NewHTTPError(http.StatusBadRequest, resp)
	}
	return nil
}

func NewValidator(e *echo.Echo) {
	e.Validator = &CustomValidator{validator: validator.New()}

}
