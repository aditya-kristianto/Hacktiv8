package repository

type UserRepo interface {
	CreateUser(user *User) error
	GetUsers() (*[]User, error)
	GetUserByID(id int) (*User, error)
}
