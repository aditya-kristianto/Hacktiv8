package user

import (
	"net/http"

	"final-project/internal/app/model"
	"final-project/internal/pkg/helper"

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

// RegisterUser godoc
// @Summary      Register
// @Description  To register
// @Tags         Users
// @Accept       json
// @Produce      json
// @Param 		 body body RegisterRequest true "Request"
// @Success      200  {object}  helper.Response
// @Success      201  {object}  helper.Response
// @Failure      400  {object}  helper.Response
// @Failure      404  {object}  helper.Response
// @Failure      405  {object}  helper.Response
// @Failure      500  {object}  helper.Response
// @Router       /users/register [post]
func (u *Controller) Register(c echo.Context) (err error) {
	req := new(RegisterRequest)
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

	err = u.service.Register(req)
	if err != nil {
		return c.JSON(http.StatusBadRequest, &helper.Response{
			Status:  http.StatusBadRequest,
			Message: "Failed register",
			Error:   err.Error(),
		})
	}

	return c.JSON(http.StatusCreated, req)
}

// Login godoc
// @Summary      Login
// @Description  To login
// @Tags         Users
// @Accept       json
// @Produce      json
// @Param 		 body body LoginRequest true "Request"
// @Success      200  {object}  helper.Response
// @Failure      400  {object}  helper.Response
// @Failure      404  {object}  helper.Response
// @Failure      405  {object}  helper.Response
// @Failure      500  {object}  helper.Response
// @Router       /users/login [post]
func (u *Controller) Login(c echo.Context) (err error) {
	req := new(LoginRequest)
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

	token, err := u.service.Login(req)
	if err != nil {
		return c.JSON(http.StatusBadRequest, &helper.Response{
			Status:  http.StatusBadRequest,
			Message: "Failed login",
			Error:   err.Error(),
		})
	}

	return c.JSON(http.StatusOK, echo.Map{
		"token": token,
	})
}

// UpdateUser godoc
// @Summary      Update user
// @Description  To update user
// @Tags         Users
// @Accept       json
// @Produce      json
// @Param 		 Authorization header string true "Bearer Token"
// @Param 		 body body UpdateUserRequest true "Request"
// @Success      200  {object}  helper.Response
// @Failure      400  {object}  helper.Response
// @Failure      401  {object}  helper.Response
// @Failure      404  {object}  helper.Response
// @Failure      405  {object}  helper.Response
// @Failure      500  {object}  helper.Response
// @Router       /users [put]
func (u *Controller) UpdateUser(c echo.Context) (err error) {
	req := new(UpdateUserRequest)
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

	userID := helper.GetUserID(c)
	data := model.User{
		Email:    req.Email,
		Username: req.Username,
	}
	result, err := u.service.UpdateUser(&userID, &data)
	if err != nil {
		return c.JSON(http.StatusBadRequest, &helper.Response{
			Status:  http.StatusBadRequest,
			Message: "Failed update user",
			Error:   err.Error(),
		})
	}

	return c.JSON(http.StatusOK, result)
}

// DeleteUser godoc
// @Summary      Delete user
// @Description  To delete user
// @Tags         Users
// @Accept       json
// @Produce      json
// @Param 		 Authorization header string true "Bearer"
// @Success      200  {object}  helper.Response
// @Failure      400  {object}  helper.Response
// @Failure      401  {object}  helper.Response
// @Failure      404  {object}  helper.Response
// @Failure      405  {object}  helper.Response
// @Failure      500  {object}  helper.Response
// @Router       /users [delete]
func (u *Controller) DeleteUser(c echo.Context) (err error) {
	userID := helper.GetUserID(c)

	u.service.DeleteUser(&userID)
	if err != nil {
		return c.JSON(http.StatusBadRequest, &helper.Response{
			Status:  http.StatusBadRequest,
			Message: "Failed delete user",
			Error:   err.Error(),
		})
	}

	return c.JSON(http.StatusOK, helper.Response{
		Message: "Your account has been successfully deleted",
	})
}
