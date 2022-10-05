package main

import (
	"net/http"
	"time"

	"github.com/go-playground/validator"
	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
	"github.com/swaggo/echo-swagger"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"

	_ "assignment-2/docs"
)

type (
	Order struct {
		ID           uuid.UUID `gorm:"primaryKey;type:uuid;default:gen_random_uuid()"`
		CustomerName string
		OrderedAt    time.Time
		Items        []Item
	}

	Item struct {
		ID          uuid.UUID `gorm:"primaryKey;type:uuid;default:gen_random_uuid()"`
		ItemCode    string
		Description string
		Quantity    int
		OrderId     uuid.UUID
		// Order       Order `gorm:"foreignKey:order_id;references:id;constraint:OnUpdate:CASCADE,OnDelete:CASCADE"`
	}

	OrderRequest struct {
		OrderedAt    time.Time     `json:"orderedAt" validate:"required" example:"2019-11-09T21:21:46+00:00"`
		CustomerName string        `json:"customerName" validate:"required" example:"Tom Jerry"`
		Items        []ItemRequest `json:"items" validate:"required"`
	}

	ItemRequest struct {
		LineItemId  int    `json:"lineItemId" validate:"omitempty" example:"1"`
		ItemCode    string `json:"itemCode" validate:"required" example:"123"`
		Description string `json:"description" validate:"required" example:"IPhone 10X"`
		Quantity    int    `json:"quantity" validate:"required" example:"1"`
	}

	Response struct {
		Status  int         `json:"status"`
		Message string      `json:"message"`
		Error   interface{} `json:"error,omitempty"`
		Payload interface{} `json:"payload,omitempty"`
	}

	CustomValidator struct {
		validator *validator.Validate
	}
)

var db *gorm.DB

// @title Swagger Assignment 2 API
// @version 1.0
// @description This is a sample server Petstore server.
// @termsOfService http://swagger.io/terms/

// @contact.name API Support
// @contact.url http://www.swagger.io/support
// @contact.email support@swagger.io

// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html

// @host localhost:1323
// @BasePath /
// @schemes http
func main() {
	dsn := "host=localhost user=aditya.kristianto@mncgroup.com password=hacktiv8 dbname=orders_by port=5432 sslmode=disable TimeZone=Asia/Jakarta"
	var err error
	db, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	// Migrate the schema
	db.AutoMigrate(&Item{})
	db.AutoMigrate(&Order{})

	e := echo.New()
	e.Validator = &CustomValidator{validator: validator.New()}

	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, World!")
	})
	e.POST("/orders", createOrder)
	e.GET("/orders", getOrders)
	e.PUT("/orders/:id", updateOrder)
	e.DELETE("/orders/:id", deleteOrder)
	e.GET("/swagger/*", echoSwagger.WrapHandler)
	e.Logger.Fatal(e.Start(":1323"))
}

func (cv *CustomValidator) Validate(i interface{}) error {
	if err := cv.validator.Struct(i); err != nil {
		// Optionally, you could return the error to give each route more control over the status code
		resp := new(Response)
		resp.Status = http.StatusBadRequest
		resp.Message = "Bad Request"
		resp.Error = err.Error()
		return echo.NewHTTPError(http.StatusBadRequest, resp)
	}
	return nil
}

// CreateOrder godoc
// @Summary      Create order
// @Description  Create order
// @Tags         Order
// @Accept       json
// @Produce      json
// @Param 		 body body OrderRequest true "Order Request"
// @Success      200  {object}  Response
// @Failure      400  {object}  Response
// @Failure      405  {object}  Response
// @Failure      500  {object}  Response
// @Router       /orders [post]
func createOrder(c echo.Context) (err error) {
	req := new(OrderRequest)
	if err = c.Bind(req); err != nil {
		resp := new(Response)
		resp.Status = http.StatusBadRequest
		resp.Message = "Bad Request"
		resp.Error = err.Error()
		return echo.NewHTTPError(http.StatusBadRequest, resp)
	}
	if err = c.Validate(req); err != nil {
		return err
	}

	err = db.Transaction(func(tx *gorm.DB) error {
		order := &Order{
			CustomerName: req.CustomerName,
			OrderedAt:    req.OrderedAt,
		}
		db.Create(order)

		for _, item := range req.Items {
			db.Create(&Item{
				ItemCode:    item.ItemCode,
				Description: item.Description,
				Quantity:    item.Quantity,
				OrderId:     order.ID,
			})
		}

		return nil
	})
	if err != nil {
		return c.JSON(http.StatusInternalServerError, &Response{
			Status:  http.StatusInternalServerError,
			Message: "Failed create order",
			Error:   err.Error(),
		})
	}

	return c.JSON(http.StatusOK, &Response{
		Status:  http.StatusOK,
		Message: "Success create order",
	})
}

// GetOrder godoc
// @Summary      Get order
// @Description  Get order
// @Tags         Order
// @Produce      json
// @Success      200  {object}  Response
// @Failure      400  {object}  Response
// @Failure      405  {object}  Response
// @Failure      500  {object}  Response
// @Router       /orders [get]
func getOrders(c echo.Context) error {
	var orders []Order
	err := db.Model(&Order{}).Preload("Items").Find(&orders).Error
	if err != nil {
		return c.JSON(http.StatusInternalServerError, &Response{
			Status:  http.StatusInternalServerError,
			Message: "Failed get order",
			Error:   err.Error(),
		})
	}

	return c.JSON(http.StatusOK, &Response{
		Status:  http.StatusOK,
		Message: "Success get order",
		Payload: orders,
	})
}

// UpdateOrder godoc
// @Summary      Update order
// @Description  Update order
// @Tags         Order
// @Accept       json
// @Produce      json
// @Param        id path string true "Order ID"
// @Param 		 body body OrderRequest true "Order Request"
// @Success      200  {object}  Response
// @Failure      400  {object}  Response
// @Failure      405  {object}  Response
// @Failure      500  {object}  Response
// @Router       /orders/{id} [put]
func updateOrder(c echo.Context) error {
	orderID, err := uuid.Parse(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, &Response{
			Status:  http.StatusBadRequest,
			Message: "Bad Request",
			Error:   err.Error(),
		})
	}

	req := new(OrderRequest)
	if err = c.Bind(req); err != nil {
		resp := new(Response)
		resp.Status = http.StatusBadRequest
		resp.Message = "Bad Request"
		resp.Error = err.Error()
		return echo.NewHTTPError(http.StatusBadRequest, resp)
	}
	if err = c.Validate(req); err != nil {
		return err
	}

	var items []Item
	err = db.Model(&Item{}).Order("id asc").Where("order_id = ?", orderID).Find(&items).Error
	if err != nil {
		return c.JSON(http.StatusInternalServerError, &Response{
			Status:  http.StatusInternalServerError,
			Message: "Failed get order",
			Error:   err.Error(),
		})
	}

	err = db.Model(Item{}).Where("id = ?", orderID).Updates(Order{
		CustomerName: req.CustomerName,
		OrderedAt:    req.OrderedAt,
	}).Error
	if err != nil {
		return err
	}

	err = db.Transaction(func(tx *gorm.DB) error {
		err = db.Model(Order{}).Where("id = ?", orderID).Updates(Order{
			CustomerName: req.CustomerName,
			OrderedAt:    req.OrderedAt,
		}).Error
		if err != nil {
			return err
		}

		for _, item := range req.Items {
			if item.LineItemId != 0 && item.LineItemId <= len(items) {
				err = db.Model(Item{}).Where("id = ?", items[item.LineItemId-1].ID).Updates(Item{
					ItemCode:    item.ItemCode,
					Description: item.Description,
					Quantity:    item.Quantity,
				}).Error
				if err != nil {
					return err
				}
			}
		}

		return nil
	})
	if err != nil {
		return c.JSON(http.StatusInternalServerError, &Response{
			Status:  http.StatusInternalServerError,
			Message: "Failed update order",
			Error:   err.Error(),
		})
	}

	return c.JSON(http.StatusOK, &Response{
		Status:  http.StatusOK,
		Message: "Success update order",
	})
}

// DeleteOrder godoc
// @Summary      Delete order
// @Description  Delete order
// @Tags         Order
// @Accept       json
// @Produce      json
// @Param        id path string true "Order ID"
// @Success      200  {object}  Response
// @Failure      400  {object}  Response
// @Failure      405  {object}  Response
// @Failure      500  {object}  Response
// @Router       /orders/{id} [delete]
func deleteOrder(c echo.Context) error {
	orderID, err := uuid.Parse(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, &Response{
			Status:  http.StatusBadRequest,
			Message: "Bad Request",
			Error:   err.Error(),
		})
	}

	err = db.Transaction(func(tx *gorm.DB) error {
		err = db.Where("order_id = ?", orderID).Delete(&Item{}).Error
		if err != nil {
			return err
		}

		err = db.Where("id = ?", orderID).Delete(&Order{}).Error
		if err != nil {
			return err
		}

		return nil
	})
	if err != nil {
		return c.JSON(http.StatusInternalServerError, &Response{
			Status:  http.StatusInternalServerError,
			Message: "Failed delete order",
			Error:   err.Error(),
		})
	}

	return c.JSON(http.StatusOK, &Response{
		Status:  http.StatusOK,
		Message: "Success delete order",
	})
}
