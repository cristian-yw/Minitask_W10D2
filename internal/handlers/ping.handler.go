package handlers

import (
	"net/http"

	"github.com/cristian-yw/Minitask_W10D2/internal/models"
	"github.com/cristian-yw/Minitask_W10D2/internal/utils"
	"github.com/gin-gonic/gin"
)

type PingHandler struct{}

func NewPingHandler() *PingHandler {
	return &PingHandler{}
}
func (p *PingHandler) GetPing(ctx *gin.Context) {
	id := ctx.GetHeader("id")
	contenttype := ctx.GetHeader("Content-Type")
	ctx.JSON(http.StatusOK, gin.H{
		"message": "pong",
		"id":      id,
		"content": contenttype,
	})
}
func (p *PingHandler) GetPingWithParam(ctx *gin.Context) {
	id := ctx.Param("id")
	param2 := ctx.Param("param2")
	q := ctx.Query("q")
	ctx.JSON(http.StatusOK, gin.H{
		"id":     id,
		"param2": param2,
		"q":      q,
	})
}
func (p *PingHandler) PostPing(ctx *gin.Context) {
	body := models.Body{}
	if err := ctx.ShouldBind(&body); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"Message": err.Error(),
			"succes":  false,
		})
		return
	}
	if err := utils.ValidatePost(body); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{
		"success": true,
		"body":    body,
	})
}
func (p *PingHandler) PatchPing(ctx *gin.Context) {
	Body := models.Body{}
	if err := ctx.ShouldBind(&Body); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"Message": err.Error(),
			"succes":  false,
		})
		return
	}
	if err := utils.ValidatePost(Body); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{
		"success": true,
		"updated": Body,
	})
}
