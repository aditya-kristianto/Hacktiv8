package socialmedia

import (
	"final-project/internal/pkg/helper"

	"net/http"

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

// CreateSocialmedias godoc
// @Summary      Create social media
// @Description  To create social media
// @Tags         SocialMedias
// @Accept       json
// @Produce      json
// @Param 		 Authorization header string true "Bearer"
// @Param 		 body body SocialMediaRequest true "Request"
// @Success      200  {object}  helper.Response
// @Success      201  {object}  helper.Response
// @Failure      400  {object}  helper.Response
// @Failure      401  {object}  helper.Response
// @Failure      404  {object}  helper.Response
// @Failure      405  {object}  helper.Response
// @Failure      500  {object}  helper.Response
// @Router       /socialmedias [post]
func (s *Controller) CreateSocialmedias(c echo.Context) (err error) {
	userID := helper.GetUserID(c)
	if userID.String() == "" {
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

	socialmedia, err := s.service.CreateSocialmedia(&userID, req)
	if err != nil {
		return c.JSON(http.StatusBadRequest, &helper.Response{
			Status:  http.StatusBadRequest,
			Message: "Failed create social media",
			Error:   err.Error(),
		})
	}

	resp := &CreateResponse{
		ID:             socialmedia.ID,
		Name:           socialmedia.Name,
		SocialMediaURL: socialmedia.SocialMediaURL,
		UserID:         socialmedia.UserID,
		CreatedAt:      socialmedia.CreatedAt,
	}
	return c.JSON(http.StatusCreated, resp)
}

// GetSocialmedias godoc
// @Summary      Get social media
// @Description  To get social media
// @Tags         SocialMedias
// @Accept       json
// @Produce      json
// @Param 		 Authorization header string true "Bearer"
// @Success      200  {object}  helper.Response
// @Failure      400  {object}  helper.Response
// @Failure      401  {object}  helper.Response
// @Failure      404  {object}  helper.Response
// @Failure      405  {object}  helper.Response
// @Failure      500  {object}  helper.Response
// @Router       /socialmedias [get]
func (s *Controller) GetSocialmedias(c echo.Context) (err error) {
	userID := helper.GetUserID(c)
	if err != nil {
		return c.JSON(http.StatusBadRequest, &helper.Response{
			Status:  http.StatusBadRequest,
			Message: "Failed get social media",
			Error:   err.Error(),
		})
	}

	socialmedias, err := s.service.GetSocialmedia(&userID)
	if err != nil {
		return c.JSON(http.StatusBadRequest, &helper.Response{
			Status:  http.StatusBadRequest,
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
// @Success      200  {object}  helper.Response
// @Failure      400  {object}  helper.Response
// @Failure      401  {object}  helper.Response
// @Failure      404  {object}  helper.Response
// @Failure      405  {object}  helper.Response
// @Failure      500  {object}  helper.Response
// @Router       /socialmedias/{socialMediaId} [put]
func (s *Controller) UpdateSocialmedias(c echo.Context) (err error) {
	socialmediaId, err := uuid.Parse(c.Param("socialMediaId"))
	if err != nil {
		return err
	}

	userID := helper.GetUserID(c)

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

	socialmedia, err := s.service.UpdateSocialmedia(&socialmediaId, &userID, req)
	if err != nil {
		return c.JSON(http.StatusBadRequest, &helper.Response{
			Status:  http.StatusBadRequest,
			Message: "Failed update social media",
			Error:   err.Error(),
		})
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
// @Success      200  {object}  helper.Response
// @Failure      400  {object}  helper.Response
// @Failure      401  {object}  helper.Response
// @Failure      404  {object}  helper.Response
// @Failure      405  {object}  helper.Response
// @Failure      500  {object}  helper.Response
// @Router       /socialmedias/{socialMediaId} [delete]
func (s *Controller) DeleteSocialmedias(c echo.Context) (err error) {
	userID := helper.GetUserID(c)

	socialmediaId, err := uuid.Parse(c.Param("socialMediaId"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, &helper.Response{
			Status:  http.StatusBadRequest,
			Message: "Failed delete social media",
			Error:   err.Error(),
		})
	}

	err = s.service.DeleteSocialmedia(&socialmediaId, &userID)
	if err != nil {
		return c.JSON(http.StatusBadRequest, &helper.Response{
			Status:  http.StatusBadRequest,
			Message: "Failed delete social media",
			Error:   err.Error(),
		})
	}

	return c.JSON(http.StatusOK, helper.Response{
		Message: "Your social media has been successfully deleted",
	})
}
