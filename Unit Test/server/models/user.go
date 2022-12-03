package models

type User struct {
	Id       int
	Email    string
	Password string
	Role     string
}

var Users = []User{}

func FindByEmail(email string) *User {
	for _, u := range Users {
		if u.Email == email {
			return &u
		}
	}

	return nil
}
