package comment

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

// CreateComments godoc
// @Summary      Create comment
// @Description  To create comment
// @Tags         Comments
// @Accept       json
// @Produce      json
// @Param 		 Authorization header string true "Bearer"
// @Param 		 body body CreateCommentRequest true "Request"
// @Success      200  {object}  helper.Response
// @Success      201  {object}  helper.Response
// @Failure      400  {object}  helper.Response
// @Failure      401  {object}  helper.Response
// @Failure      404  {object}  helper.Response
// @Failure      405  {object}  helper.Response
// @Failure      500  {object}  helper.Response
// @Router       /comments [post]
func (m *Controller) CreateComments(c echo.Context) (err error) {
	userID := helper.GetUserID(c)

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

	comment, err := m.service.CreateComment(&userID, req)
	if err != nil {
		return c.JSON(http.StatusBadRequest, &helper.Response{
			Status:  http.StatusBadRequest,
			Message: "Failed create comment",
			Error:   err.Error(),
		})
	}

	return c.JSON(http.StatusCreated, comment)
}

// GetComments godoc
// @Summary      Get comment
// @Description  To get comment
// @Tags         Comments
// @Accept       json
// @Produce      json
// @Param 		 Authorization header string true "Bearer"
// @Success      200  {object}  helper.Response
// @Failure      400  {object}  helper.Response
// @Failure      401  {object}  helper.Response
// @Failure      404  {object}  helper.Response
// @Failure      405  {object}  helper.Response
// @Failure      500  {object}  helper.Response
// @Router       /comments [get]
func (m *Controller) GetComments(c echo.Context) (err error) {
	userID := helper.GetUserID(c)

	comments, err := m.service.GetComment(&userID)
	if err != nil {
		return c.JSON(http.StatusBadRequest, &helper.Response{
			Status:  http.StatusBadRequest,
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
// @Success      200  {object}  helper.Response
// @Failure      400  {object}  helper.Response
// @Failure      401  {object}  helper.Response
// @Failure      404  {object}  helper.Response
// @Failure      405  {object}  helper.Response
// @Failure      500  {object}  helper.Response
// @Router       /comments/{commentId} [put]
func (m *Controller) UpdateComments(c echo.Context) (err error) {
	commentId, err := uuid.Parse(c.Param("commentId"))
	if err != nil {
		return err
	}

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

	data := model.Comment{
		Message: req.Message,
	}
	comment, err := m.service.UpdateComment(&commentId, &data)
	if err != nil {
		return c.JSON(http.StatusBadRequest, &helper.Response{
			Status:  http.StatusBadRequest,
			Message: "Failed update comment",
			Error:   err.Error(),
		})
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
// @Success      200  {object}  helper.Response
// @Failure      400  {object}  helper.Response
// @Failure      401  {object}  helper.Response
// @Failure      404  {object}  helper.Response
// @Failure      405  {object}  helper.Response
// @Failure      500  {object}  helper.Response
// @Router       /comments/{commentId} [delete]
func (m *Controller) DeleteComments(c echo.Context) (err error) {
	userID := helper.GetUserID(c)
	commentId, err := uuid.Parse(c.Param("commentId"))
	if err != nil {
		return err
	}

	err = m.service.DeleteComment(&commentId, &userID)
	if err != nil {
		return c.JSON(http.StatusBadRequest, &helper.Response{
			Status:  http.StatusBadRequest,
			Message: "Failed delete comment",
			Error:   err.Error(),
		})
	}

	return c.JSON(http.StatusOK, helper.Response{
		Message: "Your comment has been successfully deleted",
	})
}
