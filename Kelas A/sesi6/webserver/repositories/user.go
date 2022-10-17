package repositories

import (
	"fmt"
	"time"
)

type User struct {
	ID        int
	Name      string
	Address   string
	CreatedAt time.Time
	UpdatedAt time.Time
}

var Users = []User{
	{
		ID:      1,
		Name:    "Hacktiv8",
		Address: "Jakarta",
	},
}

func CreateUser(user *User) error {
	Users = append(Users, *user)
	return nil
}

func GetUsers() ([]User, error) {
	if len(Users) == 0 {
		return nil, fmt.Errorf("No Data !")
	}

	return Users, nil
}
