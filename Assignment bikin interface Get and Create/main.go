package main

import (
	"fmt"
	"sync"
)

type friend struct {
	name    string
	address string
	job     string
	reason  string
}

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

	fmt.Println(user.Nama + " berhasil didaftarkan")

	return ""
}

func (u service) GetUser() {
	for _, v := range u.users {
		fmt.Println(v.Nama)
	}
}

func main() {
	friends := []friend{
		{
			name:    "clara",
			address: "bekasi",
			job:     "developer",
			reason:  "training",
		}, {
			name:    "fiqri",
			address: "bekasi",
			job:     "developer",
			reason:  "training",
		}, {
			name:    "medy",
			address: "bekasi",
			job:     "developer",
			reason:  "training",
		}, {
			name:    "lutfi",
			address: "bekasi",
			job:     "developer",
			reason:  "training",
		}, {
			name:    "raiza",
			address: "bekasi",
			job:     "developer",
			reason:  "training",
		}, {
			name:    "ian",
			address: "bekasi",
			job:     "developer",
			reason:  "training",
		}, {
			name:    "umar",
			address: "bekasi",
			job:     "developer",
			reason:  "training",
		}, {
			name:    "fazry",
			address: "bekasi",
			job:     "developer",
			reason:  "training",
		}, {
			name:    "pandi",
			address: "bekasi",
			job:     "developer",
			reason:  "training",
		}, {
			name:    "tantut",
			address: "bekasi",
			job:     "developer",
			reason:  "training",
		},
	}

	var usernya []User

	var wg sync.WaitGroup

	for _, v := range friends {
		wg.Add(1)

		go func(v2 friend) {
			usernya = append(usernya, User{
				Nama: v2.name,
			})
			fmt.Println(usernya)
			wg.Done()
		}(v)
	}

	wg.Wait()

	var objectService = NewUserSvc(usernya)

	// objectService.Register(User{
	// 	Nama: "adit",
	// })
	// objectService.Register(User{
	// 	Nama: "ivan",
	// })
	// objectService.Register(User{
	// 	Nama: "clara",
	// })
	objectService.GetUser()
}
