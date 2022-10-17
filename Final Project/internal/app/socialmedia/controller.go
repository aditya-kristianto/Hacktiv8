package socialmedia

import (
	"final-project/internal/app/model"
	"final-project/internal/pkg/helper"
	"fmt"

	"net/http"

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

// CreateSocialmedias godoc
// @Summary      Create social media
// @Description  To create social media
// @Tags         SocialMedias
// @Accept       json
// @Produce      json
// @Param 		 Authorization header string true "Bearer"
// @Param 		 body body SocialMediaRequest true "Request"
// @Success      200  {object}  Response
// @Success      201  {object}  Response
// @Failure      400  {object}  Response
// @Failure      401  {object}  Response
// @Failure      404  {object}  Response
// @Failure      405  {object}  Response
// @Failure      500  {object}  Response
// @Router       /socialmedias [post]
func (s *Controller) CreateSocialmedias(c echo.Context) (err error) {
	UserID := helper.GetUserID(c)
	if UserID.String() == "" {
		resp := new(helper.Response)
		resp.Status = http.StatusBadRequest
		resp.Message = "Bad Request"
		resp.Error = err.Error()
		return echo.NewHTTPError(http.StatusBadRequest, resp)
	}

	req := new(SocialMediaRequest)
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

	data := &model.SocialMedia{
		Name:           req.Name,
		SocialMediaURL: req.SocialMediaURL,
		UserID:         UserID,
	}

	err = s.db.Create(data).Error
	if err != nil {
		fmt.Println("CreateSocialmedias 6")
		return c.JSON(http.StatusInternalServerError, &helper.Response{
			Status:  http.StatusInternalServerError,
			Message: "Failed create social media",
			Error:   err.Error(),
		})
	}

	return c.JSON(http.StatusCreated, data)
}

// GetSocialmedias godoc
// @Summary      Get social media
// @Description  To get social media
// @Tags         SocialMedias
// @Accept       json
// @Produce      json
// @Param 		 Authorization header string true "Bearer"
// @Success      200  {object}  Response
// @Failure      400  {object}  Response
// @Failure      401  {object}  Response
// @Failure      404  {object}  Response
// @Failure      405  {object}  Response
// @Failure      500  {object}  Response
// @Router       /socialmedias [get]
func (s *Controller) GetSocialmedias(c echo.Context) (err error) {
	UserID := helper.GetUserID(c).String()

	var socialmedias []model.SocialMedia
	err = s.db.Model(&model.SocialMedia{}).Where("user_id = ?", UserID).Find(&socialmedias).Error
	if err != nil {
		fmt.Println(err.Error())
		return c.JSON(http.StatusInternalServerError, &helper.Response{
			Status:  http.StatusInternalServerError,
			Message: "Failed get social media",
			Error:   err.Error(),
		})
	}

	return c.JSON(http.StatusOK, echo.Map{
		"social_medias": socialmedias,
	})
}

// UpdateSocialmedias godoc
// @Summary      Update social media
// @Description  To update social media
// @Tags         SocialMedias
// @Accept       json
// @Produce      json
// @Param 		 Authorization header string true "Bearer"
// @Param        socialMediaId path string true "Social Media ID"
// @Param 		 body body SocialMediaRequest true "Request"
// @Success      200  {object}  Response
// @Failure      400  {object}  Response
// @Failure      401  {object}  Response
// @Failure      404  {object}  Response
// @Failure      405  {object}  Response
// @Failure      500  {object}  Response
// @Router       /socialmedias/{socialMediaId} [put]
func (s *Controller) UpdateSocialmedias(c echo.Context) (err error) {
	socialMediaId := c.Param("socialMediaId")

	UserID := helper.GetUserID(c)

	req := new(SocialMediaRequest)
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

	var socialmedia model.SocialMedia
	err = s.db.Model(&socialmedia).Clauses(clause.Returning{}).Where("id = ? and user_id = ?", socialMediaId, UserID).Updates(model.SocialMedia{
		Name:           req.Name,
		SocialMediaURL: req.SocialMediaURL,
	}).Error
	if err != nil {
		return err
	}

	return c.JSON(http.StatusOK, socialmedia)
}

// DeleteSocialmedias godoc
// @Summary      Delete social media
// @Description  To delete social media
// @Tags         SocialMedias
// @Accept       json
// @Produce      json
// @Param 		 Authorization header string true "Bearer"
// @Param        socialMediaId path string true "Social Media ID"
// @Success      200  {object}  Response
// @Failure      400  {object}  Response
// @Failure      401  {object}  Response
// @Failure      404  {object}  Response
// @Failure      405  {object}  Response
// @Failure      500  {object}  Response
// @Router       /socialmedias/{socialMediaId} [delete]
func (s *Controller) DeleteSocialmedias(c echo.Context) (err error) {
	socialMediaId := c.Param("socialMediaId")

	var socialmedia model.SocialMedia
	err = s.db.Where("id = ?", socialMediaId).Delete(&socialmedia).Error
	if err != nil {
		return err
	}

	return c.JSON(http.StatusOK, helper.Response{
		Message: "Your social media has been successfully deleted",
	})
}
