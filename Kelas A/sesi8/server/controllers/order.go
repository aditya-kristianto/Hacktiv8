package controllers

import (
	"encoding/json"
	"net/http"
	"sesi8/server/params"
)

type Response struct {
	Message string `example:"STATUS_OK"`
	Status  int    `example:"200"`
	Payload interface{}
}

// GetOrders godoc
// @Summary Get all orders
// @Decription Get list orders
// @Tags order
// @Accept json
// @Produce json
// @Success 200 {object} Response
// @Router /orders [get]
func GetOrders(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode(Response{
		Message: "OK",
		Status:  200,
		Payload: nil,
	})
}

// CreateOrder godoc
// @Summary Create order
// @Description Create order by product id and user id
// @Tags order
// @Accept json
// @Produce json
// @Param order body params.CreateOrder true "Create Order"
// @Success 201 {object} Response
// @Router /orders [post]
func CreateOrder(w http.ResponseWriter, r *http.Request) {
	var req params.CreateOrder
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		json.NewEncoder(w).Encode(Response{
			Status:  http.StatusBadRequest,
			Message: "BAD_REQUEST",
		})
		return
	}

	// hit ke services layer

	json.NewEncoder(w).Encode(Response{
		Status:  http.StatusCreated,
		Message: "CREATE_ORDER_SUCCESS",
		Payload: req,
	})
}
