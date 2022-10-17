package model

type User struct {
	Id       int
	Email    string
	Password string
	Role     string
}

var Users = []User{}

func FindbyEmail(email string) *User {
	for _, user := range Users {
		if user.Email == email {
			return &user
		}
	}
	return nil
}
