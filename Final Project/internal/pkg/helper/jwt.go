package helper

import (
	"github.com/golang-jwt/jwt"
	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
)

type (
	JwtCustomClaims struct {
		UserID uuid.UUID `json:"user_id"`
		jwt.StandardClaims
	}
)

func GetUserID(c echo.Context) (UserID uuid.UUID) {
	token := c.Get("user").(*jwt.Token)
	claims := token.Claims.(*JwtCustomClaims)
	UserID = claims.UserID

	return
}
