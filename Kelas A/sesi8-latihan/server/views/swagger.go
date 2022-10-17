package views

import "sesi8-latihan/server/models"

type GetAllPeopleSwagger struct {
	Response
	Payload []models.Person
}
