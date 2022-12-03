package helper

import (
	"encoding/json"
	"fmt"
	"time"

	"github.com/dgrijalva/jwt-go"
)

type Token struct {
	Email   string
	Role    string
	Expired time.Time
}

var secret = "iniAdalahSecret"

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

func ValidateToken(tokString string) (*Token, error) {
	errResponse := fmt.Errorf("need signin")
	tok, err := jwt.Parse(tokString, func(t *jwt.Token) (interface{}, error) {
		if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errResponse
		}

		return []byte(secret), nil
	})

	if err != nil {
		return nil, err
	}

	claims, ok := tok.Claims.(jwt.MapClaims)
	if !ok && !tok.Valid {
		return nil, errResponse
	}

	fmt.Println(claims)

	payload, ok := claims["payload"]
	if !ok && !tok.Valid {
		return nil, errResponse
	}

	bytePayload, err := json.Marshal(payload)
	if err != nil {
		return nil, err
	}

	var token Token
	err = json.Unmarshal(bytePayload, &token)
	if err != nil {
		return nil, err
	}
	return &token, nil
}
