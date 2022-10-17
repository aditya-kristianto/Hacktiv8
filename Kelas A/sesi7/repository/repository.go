package repository

import "sesi7/model"

type UserRepository interface {
	CreateUser(user *model.User) error
	GetUsers() (*[]model.User, error)
}

type ProductRepository interface {
	CreateProduct(product *model.Product) error
	GetProducts() (*[]model.Product, error)
}
