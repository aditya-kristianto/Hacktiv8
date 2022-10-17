package helper

import "golang.org/x/crypto/bcrypt"

func GeneratePassword(password []byte) (string, error) {
	hash, err := bcrypt.GenerateFromPassword(password, bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}

	return string(hash), nil
}

func ValidatePassword(hash, pass []byte) error {
	err := bcrypt.CompareHashAndPassword(hash, pass)
	return err
}
