package service

import (
	"database/sql"
	"fmt"
	"sesi11/server/model"
	"sesi11/server/repository/postgres"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

var userRepo = postgres.NewUserRepoMock(&mock.Mock{})
var userSvc = NewUserServices(userRepo)

func TestFindOneNotFound(t *testing.T) {
	userRepo.(*postgres.UserRepoMock).Mock.On("FindById", "1").Return(nil)

	user, err := userSvc.GetSingleUserById("1")

	assert.NotNil(t, err)
	assert.Equal(t, err.Error(), sql.ErrNoRows.Error())
	assert.Nil(t, user)
}

func TestFindOneSuccess(t *testing.T) {
	userMock := model.User{
		Id:   "2",
		Name: "MNC B",
	}
	userRepo.(*postgres.UserRepoMock).Mock.On("FindById", userMock.Id).Return(userMock)

	user, err := userSvc.GetSingleUserById(userMock.Id)

	assert.Nil(t, err)
	assert.NotNil(t, user)

	fmt.Println(user, err)
	assert.Equal(t, user.Name, "MNC B")

}
