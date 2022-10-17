package comment

import (
	"net/http"

	"final-project/internal/app/model"
	"final-project/internal/pkg/helper"

	"github.com/golang-jwt/jwt"
	"github.com/google/uuid"
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
// @Failure      404  {object}  Response
// @Failure      405  {object}  Response
// @Failure      500  {object}  Response
// @Router       /comments [post]
func (m *Controller) CreateComments(c echo.Context) (err error) {
	token := c.Get("user").(*jwt.Token)
	claims := token.Claims.(*helper.JwtCustomClaims)
	UserID := claims.UserID

	req := new(CreateCommentRequest)
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

	photoID, err := uuid.Parse(req.PhotoID)
	if err != nil {
		return c.JSON(http.StatusBadRequest, &helper.Response{
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
	err = m.db.Create(data).Error
	if err != nil {
		return c.JSON(http.StatusInternalServerError, &helper.Response{
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
// @Failure      404  {object}  Response
// @Failure      405  {object}  Response
// @Failure      500  {object}  Response
// @Router       /comments [get]
func (m *Controller) GetComments(c echo.Context) (err error) {
	token := c.Get("user").(*jwt.Token)
	claims := token.Claims.(*helper.JwtCustomClaims)
	UserID := claims.UserID

	var comments []model.Comment
	err = m.db.Model(&model.Comment{}).Where("user_id = ?", UserID).Find(&comments).Error
	if err != nil {
		return c.JSON(http.StatusInternalServerError, &helper.Response{
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
// @Failure      404  {object}  Response
// @Failure      405  {object}  Response
// @Failure      500  {object}  Response
// @Router       /comments/{commentId} [put]
func (m *Controller) UpdateComments(c echo.Context) (err error) {
	commentId := c.Param("commentId")

	req := new(UpdateCommentRequest)
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

	var comment model.Comment
	err = m.db.Model(&comment).Clauses(clause.Returning{}).Where("id = ?", commentId).Updates(model.Comment{
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
// @Failure      404  {object}  Response
// @Failure      405  {object}  Response
// @Failure      500  {object}  Response
// @Router       /comments/{commentId} [delete]
func (m *Controller) DeleteComments(c echo.Context) (err error) {
	commentId := c.Param("commentId")

	var comment model.Comment
	err = m.db.Where("id = ?", commentId).Delete(&comment).Error
	if err != nil {
		return err
	}

	return c.JSON(http.StatusOK, helper.Response{
		Message: "Your comment has been successfully deleted",
	})
}
