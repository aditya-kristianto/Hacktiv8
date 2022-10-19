package main

import (
	"net/http"
	"time"

	"final-project/internal/app/model"
	"final-project/internal/pkg/helper"

	"github.com/go-playground/validator"
	"github.com/golang-jwt/jwt"
	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/swaggo/echo-swagger"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
	"gorm.io/gorm/logger"

	_ "final-project/docs"
)

type (
	RegisterRequest struct {
		Age      int    `json:"age" example:"31"`
		Email    string `json:"email" example:"aditya.kristianto@mncgroup.com"`
		Password string `json:"password" example:"hacktiv8"`
		Username string `json:"username" example:"aditya.kristianto"`
	}

	LoginRequest struct {
		Email    string `json:"email" example:"aditya.kristianto@mncgroup.com"`
		Password string `json:"password" example:"hacktiv8"`
	}

	UpdateUserRequest struct {
		Email    string `json:"email" example:"aditya.kristianto@mncgroup.com"`
		Username string `json:"username" example:"aditya.kristianto"`
	}

	PhotoRequest struct {
		Title    string `json:"title" example:"echo_golang"`
		Caption  string `json:"caption" example:"echo"`
		PhotoURL string `json:"photo_url" example:"https://cdn.labstack.com/images/echo-logo.svg"`
	}

	CreateCommentRequest struct {
		PhotoID string `json:"photo_id"`
		Message string `json:"message" example:"required"`
	}

	UpdateCommentRequest struct {
		Message string `json:"message" example:"required"`
	}

	SocialMediaRequest struct {
		Name           string `json:"name" example:"linkedin"`
		SocialMediaURL string `json:"social_media_url" example:"https://www.linkedin.com/in/aditya-kristianto"`
	}

	Response struct {
		Status  int         `json:"status,omitempty"`
		Message string      `json:"message"`
		Error   interface{} `json:"error,omitempty"`
		Payload interface{} `json:"payload,omitempty"`
	}

	CustomValidator struct {
		validator *validator.Validate
	}

	jwtCustomClaims struct {
		UserID uuid.UUID `json:"user_id"`
		jwt.StandardClaims
	}
)

var db *gorm.DB

// @title Swagger Final Project API
// @version 1.0
// @description This is a sample server Final Project server.
// @termsOfService http://swagger.io/terms/

// @contact.name API Support
// @contact.url http://www.swagger.io/support
// @contact.email kristianto.aditya@gmail.com

// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html

// @host localhost:1323
// @BasePath /
// @schemes http
func main() {
	dsn := "host=localhost user=aditya.kristianto@mncgroup.com password=hacktiv8 dbname=final_project port=5432 sslmode=disable TimeZone=Asia/Jakarta"
	var err error
	db, err = gorm.Open(postgres.Open(dsn), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})
	if err != nil {
		panic("failed to connect database")
	}

	// Migrate the schema
	db.AutoMigrate(&model.User{})
	db.AutoMigrate(&model.Photo{})
	db.AutoMigrate(&model.Comment{})
	db.AutoMigrate(&model.SocialMedia{})

	e := echo.New()

	// Middleware
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	e.Validator = &CustomValidator{validator: validator.New()}

	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, World!")
	})
	e.GET("/swagger/*", echoSwagger.WrapHandler)

	// Configure middleware with the custom claims type
	config := middleware.JWTConfig{
		Claims:     &jwtCustomClaims{},
		SigningKey: []byte("secret"),
	}

	g := e.Group("/users")
	g.POST("/register", register)
	g.POST("/login", login)
	g.PUT("", updateUser, middleware.JWTWithConfig(config))
	g.DELETE("", deleteUser, middleware.JWTWithConfig(config))

	g = e.Group("/photos")
	g.Use(middleware.JWTWithConfig(config))
	g.POST("", createPhotos)
	g.GET("", getPhotos)
	g.PUT("/:photoId", updatePhotos)
	g.DELETE("/:photoId", deletePhotos)

	g = e.Group("/comments")
	g.Use(middleware.JWTWithConfig(config))
	g.POST("", createComments)
	g.GET("", getComments)
	g.PUT("/:commentId", updateComments)
	g.DELETE("/:commentId", deleteComments)

	g = e.Group("/socialmedias")
	g.Use(middleware.JWTWithConfig(config))
	g.POST("", createSocialmedias)
	g.GET("", getSocialmedias)
	g.PUT("/:socialMediaId", updateSocialmedias)
	g.DELETE("/:socialMediaId", deleteSocialmedias)

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
// @Failure      405  {object}  Response
// @Failure      500  {object}  Response
// @Router       /users/register [post]
func register(c echo.Context) (err error) {
	req := new(RegisterRequest)
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

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, &Response{
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
	err = db.Create(data).Error
	if err != nil {
		return c.JSON(http.StatusInternalServerError, &Response{
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
// @Failure      405  {object}  Response
// @Failure      500  {object}  Response
// @Router       /users/login [post]
func login(c echo.Context) (err error) {
	req := new(LoginRequest)
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

	var user model.User
	err = db.Model(&model.User{}).Where("email = ?", req.Email).Find(&user).Error
	if err != nil {
		return c.JSON(http.StatusInternalServerError, &Response{
			Status:  http.StatusInternalServerError,
			Message: "Failed get order",
			Error:   err.Error(),
		})
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(req.Password))
	if err != nil {
		return c.JSON(http.StatusBadRequest, &Response{
			Status:  http.StatusBadRequest,
			Message: "Email or password is incorrect",
			Error:   err.Error(),
		})
	}

	claims := &jwtCustomClaims{
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
// @Failure      405  {object}  Response
// @Failure      500  {object}  Response
// @Router       /users [put]
func updateUser(c echo.Context) (err error) {
	req := new(UpdateUserRequest)
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

	UserID := helper.GetUserID(c)

	var user model.User
	err = db.Model(&user).Clauses(clause.Returning{}).Where("id = ?", UserID).Updates(model.User{
		Email:    req.Email,
		Username: req.Username,
	}).Error
	if err != nil {
		return err
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
// @Failure      405  {object}  Response
// @Failure      500  {object}  Response
// @Router       /users [delete]
func deleteUser(c echo.Context) (err error) {
	UserID := helper.GetUserID(c)

	var user model.User
	err = db.Where("id = ?", UserID).Delete(&user).Error
	if err != nil {
		return err
	}

	return c.JSON(http.StatusOK, Response{
		Message: "Your account has been successfully deleted",
	})
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
// @Failure      405  {object}  Response
// @Failure      500  {object}  Response
// @Router       /photos [post]
func createPhotos(c echo.Context) (err error) {
	req := new(PhotoRequest)
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

	data := &model.Photo{
		Title:    req.Title,
		Caption:  req.Caption,
		PhotoURL: req.PhotoURL,
	}
	err = db.Create(data).Error
	if err != nil {
		return c.JSON(http.StatusInternalServerError, &Response{
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
// @Failure      405  {object}  Response
// @Failure      500  {object}  Response
// @Router       /photos [get]
func getPhotos(c echo.Context) (err error) {
	var photos []model.Photo
	err = db.Model(&model.Photo{}).Find(&photos).Error
	if err != nil {
		return c.JSON(http.StatusInternalServerError, &Response{
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
// @Failure      405  {object}  Response
// @Failure      500  {object}  Response
// @Router       /photos/{photoId} [put]
func updatePhotos(c echo.Context) (err error) {
	photoId := c.Param("photoId")

	req := new(PhotoRequest)
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

	var photo model.Photo
	err = db.Model(&photo).Clauses(clause.Returning{}).Where("id = ?", photoId).Updates(model.Photo{
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
// @Failure      405  {object}  Response
// @Failure      500  {object}  Response
// @Router       /photos/{photoId} [delete]
func deletePhotos(c echo.Context) (err error) {
	photoId := c.Param("photoId")

	var photo model.Photo
	err = db.Where("id = ?", photoId).Delete(&photo).Error
	if err != nil {
		return err
	}

	return c.JSON(http.StatusOK, Response{
		Message: "Your photo has been successfully deleted",
	})
}

// CreateComments godoc
// @Summary      Create comment
// @Description  To create comment
// @Tags         Comments
// @Accept       json
// @Produce      json
// @Param 		 Authorization header string true "Bearer"
// @Param 		 body body CreateCommentRequest true "Request"
// @Success      200  {object}  Response
// @Success      201  {object}  Response
// @Failure      400  {object}  Response
// @Failure      401  {object}  Response
// @Failure      405  {object}  Response
// @Failure      500  {object}  Response
// @Router       /comments [post]
func createComments(c echo.Context) (err error) {
	token := c.Get("user").(*jwt.Token)
	claims := token.Claims.(*jwtCustomClaims)
	UserID := claims.UserID

	req := new(CreateCommentRequest)
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

	photoID, err := uuid.Parse(req.PhotoID)
	if err != nil {
		return c.JSON(http.StatusBadRequest, &Response{
			Status:  http.StatusBadRequest,
			Message: "Bad Request",
			Error:   err.Error(),
		})
	}

	data := &model.Comment{
		Message: req.Message,
		PhotoID: photoID,
		UserID:  UserID,
	}
	err = db.Create(data).Error
	if err != nil {
		return c.JSON(http.StatusInternalServerError, &Response{
			Status:  http.StatusInternalServerError,
			Message: "Failed register",
			Error:   err.Error(),
		})
	}

	return c.JSON(http.StatusCreated, data)
}

// GetComments godoc
// @Summary      Get comment
// @Description  To get comment
// @Tags         Comments
// @Accept       json
// @Produce      json
// @Param 		 Authorization header string true "Bearer"
// @Success      200  {object}  Response
// @Failure      400  {object}  Response
// @Failure      401  {object}  Response
// @Failure      405  {object}  Response
// @Failure      500  {object}  Response
// @Router       /comments [get]
func getComments(c echo.Context) (err error) {
	token := c.Get("user").(*jwt.Token)
	claims := token.Claims.(*jwtCustomClaims)
	UserID := claims.UserID

	var comments []model.Comment
	err = db.Model(&model.Comment{}).Where("user_id = ?", UserID).Find(&comments).Error
	if err != nil {
		return c.JSON(http.StatusInternalServerError, &Response{
			Status:  http.StatusInternalServerError,
			Message: "Failed get comments",
			Error:   err.Error(),
		})
	}

	return c.JSON(http.StatusOK, comments)
}

// UpdateComments godoc
// @Summary      Update comment
// @Description  To update comment
// @Tags         Comments
// @Accept       json
// @Produce      json
// @Param 		 Authorization header string true "Bearer"
// @Param        commentId path string true "Comment ID"
// @Param 		 body body UpdateCommentRequest true "Request"
// @Success      200  {object}  Response
// @Failure      400  {object}  Response
// @Failure      401  {object}  Response
// @Failure      405  {object}  Response
// @Failure      500  {object}  Response
// @Router       /comments/{commentId} [put]
func updateComments(c echo.Context) (err error) {
	commentId := c.Param("commentId")

	req := new(UpdateCommentRequest)
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

	var comment model.Comment
	err = db.Model(&comment).Clauses(clause.Returning{}).Where("id = ?", commentId).Updates(model.Comment{
		Message: req.Message,
	}).Error
	if err != nil {
		return err
	}

	return c.JSON(http.StatusOK, comment)
}

// DeleteComments godoc
// @Summary      Delete comment
// @Description  To delete comment
// @Tags         Comments
// @Accept       json
// @Produce      json
// @Param 		 Authorization header string true "Bearer"
// @Param        commentId path string true "Comment ID"
// @Success      200  {object}  Response
// @Failure      400  {object}  Response
// @Failure      401  {object}  Response
// @Failure      405  {object}  Response
// @Failure      500  {object}  Response
// @Router       /comments/{commentId} [delete]
func deleteComments(c echo.Context) (err error) {
	commentId := c.Param("commentId")

	var comment model.Comment
	err = db.Where("id = ?", commentId).Delete(&comment).Error
	if err != nil {
		return err
	}

	return c.JSON(http.StatusOK, Response{
		Message: "Your comment has been successfully deleted",
	})
}

// SreateSocialmedias godoc
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
// @Failure      405  {object}  Response
// @Failure      500  {object}  Response
// @Router       /socialmedias [post]
func createSocialmedias(c echo.Context) (err error) {
	token := c.Get("user").(*jwt.Token)
	claims := token.Claims.(*jwtCustomClaims)
	UserID := claims.UserID

	req := new(SocialMediaRequest)
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

	data := &model.SocialMedia{
		Name:           req.Name,
		SocialMediaURL: req.SocialMediaURL,
		UserID:         UserID,
	}
	err = db.Create(data).Error
	if err != nil {
		return c.JSON(http.StatusInternalServerError, &Response{
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
// @Failure      405  {object}  Response
// @Failure      500  {object}  Response
// @Router       /socialmedias [get]
func getSocialmedias(c echo.Context) (err error) {
	token := c.Get("user").(*jwt.Token)
	claims := token.Claims.(*jwtCustomClaims)
	UserID := claims.UserID

	var socialmedias []model.SocialMedia
	err = db.Model(&model.SocialMedia{}).Where("user_id = ?", UserID).Find(&socialmedias).Error
	if err != nil {
		return c.JSON(http.StatusInternalServerError, &Response{
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
// @Failure      405  {object}  Response
// @Failure      500  {object}  Response
// @Router       /socialmedias/{socialMediaId} [put]
func updateSocialmedias(c echo.Context) (err error) {
	socialMediaId := c.Param("socialMediaId")

	token := c.Get("user").(*jwt.Token)
	claims := token.Claims.(*jwtCustomClaims)
	UserID := claims.UserID

	req := new(SocialMediaRequest)
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

	var socialmedia model.SocialMedia
	err = db.Model(&socialmedia).Clauses(clause.Returning{}).Where("id = ? and user_id = ?", socialMediaId, UserID).Updates(model.SocialMedia{
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
// @Failure      405  {object}  Response
// @Failure      500  {object}  Response
// @Router       /socialmedias/{socialMediaId} [delete]
func deleteSocialmedias(c echo.Context) (err error) {
	socialMediaId := c.Param("socialMediaId")

	var socialmedia model.SocialMedia
	err = db.Where("id = ?", socialMediaId).Delete(&socialmedia).Error
	if err != nil {
		return err
	}

	return c.JSON(http.StatusOK, Response{
		Message: "Your social media has been successfully deleted",
	})
}
