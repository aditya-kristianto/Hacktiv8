package helper

import (
	"fmt"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestGenerateJwt(t *testing.T) {
	payload := Token{
		Email:   "a@mnc.com",
		Expired: time.Now(),
	}

	tok, err := GenerateToken(&payload)
	assert.Nil(t, err)
	assert.NotEmpty(t, tok)
	fmt.Println(tok)
}

func TestValidateJwt(t *testing.T) {
	token := "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJwYXlsb2FkIjp7IkVtYWlsIjoiYUBtbmMuY29tIiwiRXhwaXJlZCI6IjIwMjItMTAtMDVUMTk6MzU6MjEuNTEyNzg4KzA3OjAwIn19.raWDhlg5ZAsgwibC3Cjq-1OhGLHjaotdmTXS2WjXAQQ"

	payload, err := VerifyToken(token)

	assert.Nil(t, err)
	assert.NotNil(t, payload)
	assert.NotEmpty(t, payload.Email)
}
