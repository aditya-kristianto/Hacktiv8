package postgres

import (
	"fmt"
	"sesi11/server/model"
	"sesi11/server/repository"
)

type userRepo struct{}

func NewUserRepo() repository.UserRepository {
	return &userRepo{}
}

func (u *userRepo) FindById(id string) *model.User {
	query := "SELECT * FROM users WHERE id=$1"

	fmt.Println(query)
	return nil
}
