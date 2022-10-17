package controller

import (
	"encoding/json"
	"net/http"
	"swaggo/server/models"
	"swaggo/server/params"
	"swaggo/server/views"
	"time"
)

// GetOrders godoc
// @Summary Get all orders
// @Description Get all orders in detail
// @Tags order
// @Accept json
// @Produce json
// @Success 200 {object} views.GetOrdersSuccess
// @Failure 404 {object} views.GetOrdersFailNotFound
// @Failure 500 {object} views.GetOrdersFailNotFound
// @Param limit query int false "Limit"
// @Router /orders [get]
func GetOrders(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(views.Response{
		StatusCode: 200,
		Message:    "GET ORDERS SUCCESS",
		Payload: []models.Order{
			{
				Id:           1,
				CustomerName: "MNC A",
				ProductId:    10,
				UserId:       1,
				CreatedAt:    time.Now(),
			},
		},
	})
}

// CreateOrder godoc
// @Summary Create new orders
// @Description Create new orders
// @Tags order
// @Accept json
// @Produce json
// @Param order body params.CreateOrderReq true "Create order payload"
// @Param id path int true "path"
// @Param Authorization header string true "token"
// @Param X-Token header string true "X-Token"
// @Success 200 {object} views.CreateOrderSuccess
// @Failure 400 {object} views.CreateOrderFailBadRequest
// @Router /orders/{id} [post]
func CreateOrder(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var req params.CreateOrderReq

	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		json.NewEncoder(w).Encode(views.Response{
			StatusCode: 400,
			Message:    "BAD_REQUEST",
		})
		return
	}

	json.NewEncoder(w).Encode(views.Response{
		StatusCode: 200,
		Message:    "CREATE_ORDER_SUCCESS",
		Payload:    req,
	})
}
