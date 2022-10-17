package views

import (
	"swaggo/server/models"
	"swaggo/server/params"
)

type GetOrdersSuccess struct {
	StatusCode int            `json:"status" example:"200"`
	Message    string         `json:"message" example:"GET_SUCCESS"`
	Payload    []models.Order `json:"payload"`
}
type GetOrdersFailNotFound struct {
	StatusCode int    `json:"status" example:"404"`
	Message    string `json:"message" example:"GET_ORDER_NOT_FOUND"`
}

type CreateOrderSuccess struct {
	StatusCode int                   `json:"status" example:"200"`
	Message    string                `json:"message" example:"CREATED_ORDER_SUCCESS"`
	Payload    params.CreateOrderReq `json:"payload"`
}
type CreateOrderFailBadRequest struct {
	StatusCode int    `json:"status" example:"400"`
	Message    string `json:"message" example:"BAD_REQUEST"`
}
