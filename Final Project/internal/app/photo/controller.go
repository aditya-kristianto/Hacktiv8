package photo

import (
	"net/http"

	"final-project/internal/app/model"
	"final-project/internal/pkg/helper"

	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

type (
	Controller struct {
		service *Service
	}
)

func NewController(db *gorm.DB) *Controller {
	return &Controller{
		service: NewService(db),
	}
}

// CreatePhotos godoc
// @Summary      Create photo
// @Description  To create photo
// @Tags         Photos
// @Accept       json
// @Produce      json
// @Param 		 Authorization header string true "Bearer"
// @Param 		 body body PhotoRequest true "Request"
// @Success      200  {object}  helper.Response
// @Success      201  {object}  helper.Response
// @Failure      400  {object}  helper.Response
// @Failure      401  {object}  helper.Response
// @Failure      404  {object}  helper.Response
// @Failure      405  {object}  helper.Response
// @Failure      500  {object}  helper.Response
// @Router       /photos [post]
func (p *Controller) CreatePhotos(c echo.Context) (err error) {
	userID := helper.GetUserID(c)

	req := new(PhotoRequest)
	if err = c.Bind(req); err != nil {
		resp := new(helper.Response)
		resp.Status = http.StatusBadRequest
		resp.Message = "Bad Request"
		resp.Error = err.Error()
		return echo.NewHTTPError(http.StatusBadRequest, resp)
	}
	if err = c.Validate(req); err != nil {
		return err
	}

	photo, err := p.service.CreatePhoto(&userID, req)
	if err != nil {
		return c.JSON(http.StatusBadRequest, &helper.Response{
			Status:  http.StatusBadRequest,
			Message: "Failed create photo",
			Error:   err.Error(),
		})
	}

	return c.JSON(http.StatusCreated, photo)
}

// GetPhotos godoc
// @Summary      Get photo
// @Description  To get photo
// @Tags         Photos
// @Accept       json
// @Produce      json
// @Param 		 Authorization header string true "Bearer"
// @Success      200  {object}  helper.Response
// @Failure      400  {object}  helper.Response
// @Failure      401  {object}  helper.Response
// @Failure      404  {object}  helper.Response
// @Failure      405  {object}  helper.Response
// @Failure      500  {object}  helper.Response
// @Router       /photos [get]
func (p *Controller) GetPhotos(c echo.Context) (err error) {
	userID := helper.GetUserID(c)

	photos, err := p.service.GetPhoto(&userID)
	if err != nil {
		return c.JSON(http.StatusBadRequest, &helper.Response{
			Status:  http.StatusBadRequest,
			Message: "Failed get photos",
			Error:   err.Error(),
		})
	}

	return c.JSON(http.StatusOK, photos)
}

// UpdatePhotos godoc
// @Summary      Update photo
// @Description  To update photo
// @Tags         Photos
// @Accept       json
// @Produce      json
// @Param 		 Authorization header string true "Bearer"
// @Param        photoId path string true "Photo ID"
// @Param 		 body body PhotoRequest true "Request"
// @Success      200  {object}  helper.Response
// @Failure      400  {object}  helper.Response
// @Failure      401  {object}  helper.Response
// @Failure      404  {object}  helper.Response
// @Failure      405  {object}  helper.Response
// @Failure      500  {object}  helper.Response
// @Router       /photos/{photoId} [put]
func (p *Controller) UpdatePhotos(c echo.Context) (err error) {
	photoId, err := uuid.Parse(c.Param("photoId"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, &helper.Response{
			Status:  http.StatusBadRequest,
			Message: "Failed update photo",
			Error:   err.Error(),
		})
	}

	req := new(PhotoRequest)
	if err = c.Bind(req); err != nil {
		resp := new(helper.Response)
		resp.Status = http.StatusBadRequest
		resp.Message = "Bad Request"
		resp.Error = err.Error()
		return echo.NewHTTPError(http.StatusBadRequest, resp)
	}
	if err = c.Validate(req); err != nil {
		return err
	}

	data := model.Photo{
		Title:    req.Title,
		Caption:  req.Caption,
		PhotoURL: req.PhotoURL,
	}
	photo, err := p.service.UpdatePhoto(&photoId, &data)
	if err != nil {
		return c.JSON(http.StatusBadRequest, &helper.Response{
			Status:  http.StatusBadRequest,
			Message: "Failed update photo",
			Error:   err.Error(),
		})
	}

	return c.JSON(http.StatusOK, photo)
}

// DeletePhotos godoc
// @Summary      Delete photo
// @Description  To delete photo
// @Tags         Photos
// @Accept       json
// @Produce      json
// @Param 		 Authorization header string true "Bearer"
// @Param        photoId path string true "Photo ID"
// @Success      200  {object}  helper.Response
// @Failure      400  {object}  helper.Response
// @Failure      401  {object}  helper.Response
// @Failure      404  {object}  helper.Response
// @Failure      405  {object}  helper.Response
// @Failure      500  {object}  helper.Response
// @Router       /photos/{photoId} [delete]
func (p *Controller) DeletePhotos(c echo.Context) (err error) {
	photoId, err := uuid.Parse(c.Param("photoId"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, &helper.Response{
			Status:  http.StatusBadRequest,
			Message: "Failed delete photo",
			Error:   err.Error(),
		})
	}

	err = p.service.DeletePhoto(&photoId)
	if err != nil {
		return c.JSON(http.StatusBadRequest, &helper.Response{
			Status:  http.StatusBadRequest,
			Message: "Failed delete photo",
			Error:   err.Error(),
		})
	}

	return c.JSON(http.StatusOK, helper.Response{
		Message: "Your photo has been successfully deleted",
	})
}
