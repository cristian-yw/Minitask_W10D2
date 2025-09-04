package handlers

import (
	"net/http"
	"strconv"

	"github.com/cristian-yw/Minitask_W10D2/internal/models"
	"github.com/cristian-yw/Minitask_W10D2/internal/repositories"
	"github.com/gin-gonic/gin"
)

type UserHandler struct {
	sr *repositories.UserRepository
}

func NewUserHandler(sr *repositories.UserRepository) *UserHandler {
	return &UserHandler{sr: sr}
}

func (u *UserHandler) GetUsers(ctx *gin.Context) {
	page, err := strconv.Atoi(ctx.Query("page"))
	if err != nil {
		page = 1
	}
	limit := 2
	offset := (page - 1) * limit

	users, err := u.sr.GetAllUsers(ctx.Request.Context(), offset, limit)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"success": false,
			"data":    []any{},
		})
		return
	}
	if len(users) == 0 {
		ctx.JSON(http.StatusNotFound, gin.H{
			"success": true,
			"data":    []any{},
			"page":    page,
		})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{
		"success": true,
		"data":    users,
		"page":    page,
	})
}

func (u *UserHandler) UpdateUser(ctx *gin.Context) {
	idParam := ctx.Param("id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{
			"success": false,
			"error":   "Invalid user id",
		})
		return
	}
	var body models.User
	if err := ctx.ShouldBind(&body); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"success": false,
			"error":   err.Error(),
		})
		return
	}

	updatedUser, err := u.sr.UpdateUser(ctx.Request.Context(), id, body)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"success": false,
			"error":   err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"success": true,
		"message": "User updated successfully",
		"data":    updatedUser,
	})
}
