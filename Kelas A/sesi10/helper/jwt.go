package helper

import (
	"fmt"
	"time"

	"github.com/dgrijalva/jwt-go"
)

var secret = "secret"
var timeout = 1 * time.Minute

func GenerateToken(email, role string) (string, error) {
	payload := jwt.MapClaims{
		"email": email,
		"role":  role,
		"exp":   time.Now().Add(timeout),
	}

	jwtToken := jwt.NewWithClaims(jwt.SigningMethodHS256, payload)

	signed, err := jwtToken.SignedString([]byte(secret))
	if err != nil {
		return "", err
	}

	return signed, nil
}

func ValidateToken(tokenString string) (map[string]interface{}, error) {
	errResp := fmt.Errorf("need signin")
	token, err := jwt.Parse(tokenString, func(t *jwt.Token) (interface{}, error) {
		if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errResp
		}

		return []byte(secret), nil
	})

	if err != nil {
		return nil, err
	}

	if _, ok := token.Claims.(jwt.MapClaims); !ok && !token.Valid {
		return nil, errResp
	}

	var payload = map[string]interface{}{}
	claims := token.Claims.(jwt.MapClaims)
	payload["email"] = claims["email"]
	payload["role"] = claims["role"]

	exp := fmt.Sprintf("%v", claims["exp"])

	now := time.Now()
	expTime, _ := time.Parse(time.RFC3339, exp)

	if !now.Before(expTime) {
		return nil, fmt.Errorf("expired")
	}

	return payload, nil
}
