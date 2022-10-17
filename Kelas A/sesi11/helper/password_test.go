package helper

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGeneratePassword(t *testing.T) {
	pass := []byte("Hello")

	password, err := GeneratePassword(pass)
	fmt.Println(password)
	assert.Nil(t, err)
}

func TestValidatePassword(t *testing.T) {
	hash := []byte("$2a$10$KWsU1QBmAjZnG5N6J/z0D.yfCN.3xTgncSCU3Z0lfna855QuTNbki")
	pass := []byte("Hello")

	err := ValidatePassword(hash, pass)

	assert.Nil(t, err)
}

func TestInvalidatePassword(t *testing.T) {
	hash := []byte("$2a$10$KWsU1QBmAjZnG5N6J/z0D.yfCN.3xTgncSCU3Z0lfna855QuTNbk")
	pass := []byte("Hello")

	err := ValidatePassword(hash, pass)

	assert.NotNil(t, err)
}
