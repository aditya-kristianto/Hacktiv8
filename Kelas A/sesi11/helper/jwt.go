package helper

import (
	"encoding/json"
	"fmt"
	"time"

	"github.com/dgrijalva/jwt-go"
)

type Token struct {
	Email   string
	Expired time.Time
}

const secret = "iniSecret"

func GenerateToken(payload *Token) (string, error) {
	claims := jwt.MapClaims{
		"payload": payload,
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tok, err := token.SignedString([]byte(secret))
	if err != nil {
		return "", err
	}

	return tok, nil
}

func VerifyToken(tokenString string) (*Token, error) {
	errResp := fmt.Errorf("need signin")
	tok, err := jwt.Parse(tokenString, func(t *jwt.Token) (interface{}, error) {
		if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errResp
		}

		return []byte(secret), nil
	})

	if err != nil {
		return nil, err
	}

	if _, ok := tok.Claims.(jwt.MapClaims); !ok && !tok.Valid {
		return nil, errResp
	}

	claims := tok.Claims.(jwt.MapClaims)

	payloadByte, err := json.Marshal(claims["payload"])
	if err != nil {
		return nil, err
	}

	var payload Token

	err = json.Unmarshal(payloadByte, &payload)
	if err != nil {
		return nil, err
	}

	return &payload, nil

}
