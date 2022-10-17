package postgres

import (
	"sesi11/server/model"
	"sesi11/server/repository"

	"github.com/stretchr/testify/mock"
)

type UserRepoMock struct {
	Mock *mock.Mock
}

func NewUserRepoMock(mock *mock.Mock) repository.UserRepository {
	return &UserRepoMock{
		Mock: mock,
	}
}

func (u *UserRepoMock) FindById(id string) *model.User {
	arguments := u.Mock.Called(id)

	if arguments.Get(0) == nil {
		return nil
	}

	user := arguments.Get(0).(model.User)
	return &user
}
