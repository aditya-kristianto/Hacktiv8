package photo

import (
	"net/http"

	"final-project/internal/app/model"
	"final-project/internal/pkg/helper"

	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type (
	Controller struct {
		db *gorm.DB
	}
)

func NewController(db *gorm.DB) *Controller {
	return &Controller{
		db: db,
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
// @Success      200  {object}  Response
// @Failure      400  {object}  Response
// @Failure      401  {object}  Response
// @Failure      404  {object}  Response
// @Failure      405  {object}  Response
// @Failure      500  {object}  Response
// @Router       /photos [post]
func (p *Controller) CreatePhotos(c echo.Context) (err error) {
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

	data := &model.Photo{
		Title:    req.Title,
		Caption:  req.Caption,
		PhotoURL: req.PhotoURL,
	}
	err = p.db.Create(data).Error
	if err != nil {
		return c.JSON(http.StatusInternalServerError, &helper.Response{
			Status:  http.StatusInternalServerError,
			Message: "Failed register",
			Error:   err.Error(),
		})
	}

	return c.JSON(http.StatusCreated, data)
}

// GetPhotos godoc
// @Summary      Get photo
// @Description  To get photo
// @Tags         Photos
// @Accept       json
// @Produce      json
// @Param 		 Authorization header string true "Bearer"
// @Success      200  {object}  Response
// @Failure      400  {object}  Response
// @Failure      401  {object}  Response
// @Failure      404  {object}  Response
// @Failure      405  {object}  Response
// @Failure      500  {object}  Response
// @Router       /photos [get]
func (p *Controller) GetPhotos(c echo.Context) (err error) {
	var photos []model.Photo
	err = p.db.Model(&model.Photo{}).Find(&photos).Error
	if err != nil {
		return c.JSON(http.StatusInternalServerError, &helper.Response{
			Status:  http.StatusInternalServerError,
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
// @Success      200  {object}  Response
// @Failure      400  {object}  Response
// @Failure      401  {object}  Response
// @Failure      404  {object}  Response
// @Failure      405  {object}  Response
// @Failure      500  {object}  Response
// @Router       /photos/{photoId} [put]
func (p *Controller) UpdatePhotos(c echo.Context) (err error) {
	photoId := c.Param("photoId")

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

	var photo model.Photo
	err = p.db.Model(&photo).Clauses(clause.Returning{}).Where("id = ?", photoId).Updates(model.Photo{
		Title:    req.Title,
		Caption:  req.Caption,
		PhotoURL: req.PhotoURL,
	}).Error
	if err != nil {
		return err
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
// @Success      200  {object}  Response
// @Failure      400  {object}  Response
// @Failure      401  {object}  Response
// @Failure      404  {object}  Response
// @Failure      405  {object}  Response
// @Failure      500  {object}  Response
// @Router       /photos/{photoId} [delete]
func (p *Controller) DeletePhotos(c echo.Context) (err error) {
	photoId := c.Param("photoId")

	var photo model.Photo
	err = p.db.Where("id = ?", photoId).Delete(&photo).Error
	if err != nil {
		return err
	}

	return c.JSON(http.StatusOK, helper.Response{
		Message: "Your photo has been successfully deleted",
	})
}
