package routers

import (
	"github.com/cristian-yw/Minitask_W10D2/internal/handlers"
	"github.com/cristian-yw/Minitask_W10D2/internal/repositories"
	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v5/pgxpool"
)

func InitUserRouter(router *gin.Engine, db *pgxpool.Pool) {
	userRouter := router.Group("/users")
	ur := repositories.NewUserRepository(db)
	uh := handlers.NewUserHandler(ur)
	userRouter.GET("/", uh.GetUsers)
	userRouter.PATCH("/:id", uh.UpdateUser)
}
