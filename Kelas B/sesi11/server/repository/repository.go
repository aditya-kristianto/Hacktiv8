package repository

import "sesi11/server/model"

type UserRepository interface {
	FindById(id string) *model.User
}
