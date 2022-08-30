package main

import (
	"fmt"
)

type userSvc interface {
	Register(user User) string
	GetUser()
}

func NewUserSvc(u []User) userSvc {
	return &service{users: u}
}

type User struct {
	Nama string
}

type service struct {
	users []User
}

func (u service) Register(user User) string {
	u.users = append(u.users, User{
		Nama: user.Nama,
	})

	return ""
}

func (u service) GetUser() {
	for _, v := range u.users {
		fmt.Println(v.Nama)
	}
}

func main() {
	var usernya []User
	usernya = append(usernya, User{
		Nama: "adit",
	})
	usernya = append(usernya, User{
		Nama: "ivan",
	})
	usernya = append(usernya, User{
		Nama: "clara",
	})
	var objectService = NewUserSvc(usernya)

	objectService.GetUser()
}
