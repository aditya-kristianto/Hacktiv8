package models

import "time"

type Order struct {
	Id           int       `example:"1"`
	CustomerName string    `example:"MNC B"`
	ProductId    int       `example:"10"`
	UserId       int       `example:"1"`
	CreatedAt    time.Time `example:"time"`
}
