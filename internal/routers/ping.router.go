package routers

import (
	"github.com/cristian-yw/Minitask_W10D2/internal/handlers"
	"github.com/gin-gonic/gin"
)

func InitpingRouter(router *gin.Engine) {
	pingRouter := router.Group("/ping")
	ph := handlers.NewPingHandler()
	pingRouter.GET("/", ph.GetPing)
	pingRouter.GET("/:id/:param2", ph.GetPingWithParam)
	pingRouter.POST("/", ph.PostPing)
	pingRouter.PATCH("/", ph.PatchPing)
}
