package controllers

import (
	"net/http"
	"sesi8-latihan/server/models"
	"sesi8-latihan/server/views"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type PersonController struct {
	db *gorm.DB
}

func NewPersonController(db *gorm.DB) *PersonController {
	return &PersonController{
		db: db,
	}
}

// GetPeople godoc
// @Summary Get all people
// @Decription Get list people
// @Tags person
// @Accept json
// @Produce json
// @Success 200 {object} views.GetAllPeopleSwagger
// @Router /peope [get]
func (p *PersonController) GetPeople(ctx *gin.Context) {
	var people []models.Person

	err := p.db.Find(&people).Error
	if err != nil {
		WriteJsonResponse(ctx, &views.Response{
			Status:  http.StatusInternalServerError,
			Message: "GET_PEOPLE_FAIL",
			Error:   err.Error(),
		})
		return
	}
	WriteJsonResponse(ctx, &views.Response{
		Status:  http.StatusOK,
		Message: "GET_PEOPLE_SUCCESS",
		Payload: people,
	})

}
