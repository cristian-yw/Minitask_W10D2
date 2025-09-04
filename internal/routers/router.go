package routers

import (
	"net/http"

	"github.com/cristian-yw/Minitask_W10D2/internal/middlewares"
	"github.com/cristian-yw/Minitask_W10D2/internal/models"
	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v5/pgxpool"
)

func InitRouter(db *pgxpool.Pool) *gin.Engine {
	router := gin.Default()
	router.Use(middlewares.MyLogger)
	router.Use(middlewares.CORSMiddleware)
	InitpingRouter(router)
	InitUserRouter(router, db)

	router.NoRoute(func(ctx *gin.Context) {
		ctx.JSON(http.StatusNotFound, models.Response{
			Message: "Rute Salah",
			Status:  "Rute tidak ditemukan",
		})
	})
	return router
}
