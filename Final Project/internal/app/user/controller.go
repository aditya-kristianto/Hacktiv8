package user

import (
	"net/http"
	"time"

	"final-project/internal/app/model"
	"final-project/internal/pkg/helper"

	"github.com/golang-jwt/jwt"
	"github.com/labstack/echo/v4"
	"golang.org/x/crypto/bcrypt"
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

// RegisterUser godoc
// @Summary      Register
// @Description  To register
// @Tags         Users
// @Accept       json
// @Produce      json
// @Param 		 body body RegisterRequest true "Request"
// @Success      200  {object}  Response
// @Success      201  {object}  Response
// @Failure      400  {object}  Response
// @Failure      404  {object}  Response
// @Failure      405  {object}  Response
// @Failure      500  {object}  Response
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

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, &helper.Response{
			Status:  http.StatusInternalServerError,
			Message: "Failed register",
			Error:   err.Error(),
		})
	}

	data := &model.User{
		Username: req.Username,
		Email:    req.Email,
		Password: string(hashedPassword),
		Age:      req.Age,
	}
	err = u.db.Create(data).Error
	if err != nil {
		return c.JSON(http.StatusInternalServerError, &helper.Response{
			Status:  http.StatusInternalServerError,
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
// @Success      200  {object}  Response
// @Failure      400  {object}  Response
// @Failure      404  {object}  Response
// @Failure      405  {object}  Response
// @Failure      500  {object}  Response
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

	var user model.User
	err = u.db.Model(&model.User{}).Where("email = ?", req.Email).Find(&user).Error
	if err != nil {
		return c.JSON(http.StatusInternalServerError, &helper.Response{
			Status:  http.StatusInternalServerError,
			Message: "Failed get order",
			Error:   err.Error(),
		})
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(req.Password))
	if err != nil {
		return c.JSON(http.StatusBadRequest, &helper.Response{
			Status:  http.StatusBadRequest,
			Message: "Email or password is incorrect",
			Error:   err.Error(),
		})
	}

	claims := &helper.JwtCustomClaims{
		user.ID,
		jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Hour * 72).Unix(),
		},
	}

	// Create token with claims
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	// Generate encoded token and send it as response.
	t, err := token.SignedString([]byte("secret"))
	if err != nil {
		return err
	}

	return c.JSON(http.StatusOK, echo.Map{
		"token": t,
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
// @Success      200  {object}  Response
// @Failure      400  {object}  Response
// @Failure      401  {object}  Response
// @Failure      404  {object}  Response
// @Failure      405  {object}  Response
// @Failure      500  {object}  Response
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

	UserID := helper.GetUserID(c)

	var user model.User
	err = u.db.Model(&user).Clauses(clause.Returning{}).Where("id = ?", UserID).Updates(model.User{
		Email:    req.Email,
		Username: req.Username,
	}).Error
	if err != nil {
		resp := new(helper.Response)
		resp.Status = http.StatusBadRequest
		resp.Message = "Bad Request"
		resp.Error = err.Error()
		return echo.NewHTTPError(http.StatusBadRequest, resp)
	}

	return c.JSON(http.StatusOK, user)
}

// DeleteUser godoc
// @Summary      Delete user
// @Description  To delete user
// @Tags         Users
// @Accept       json
// @Produce      json
// @Param 		 Authorization header string true "Bearer"
// @Success      200  {object}  Response
// @Failure      400  {object}  Response
// @Failure      401  {object}  Response
// @Failure      404  {object}  Response
// @Failure      405  {object}  Response
// @Failure      500  {object}  Response
// @Router       /users [delete]
func (u *Controller) DeleteUser(c echo.Context) (err error) {
	UserID := helper.GetUserID(c)

	var user model.User
	err = u.db.Where("id = ?", UserID).Delete(&user).Error
	if err != nil {
		return err
	}

	return c.JSON(http.StatusOK, helper.Response{
		Message: "Your account has been successfully deleted",
	})
}
