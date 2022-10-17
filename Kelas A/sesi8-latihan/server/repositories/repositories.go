package repositories

import "sesi8-latihan/server/models"

type PersonRepo interface {
	GetPeople() (*[]models.Person, error)
}
