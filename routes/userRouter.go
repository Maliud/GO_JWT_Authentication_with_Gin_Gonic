package routes

import (
	"github.com/Maliud/GO_JWT_Authentication_with_Gin_Gonic/controllers"
	"github.com/Maliud/GO_JWT_Authentication_with_Gin_Gonic/middleware"
	"github.com/gin-gonic/gin"
)

func UserRoutes(incomingRoutes *gin.Engine) {
	incomingRoutes.Use(middleware.Authenticate())
	incomingRoutes.GET("/users", controllers.GetUsers())
	incomingRoutes.GET("/users/:user_id", controllers.GetUser())
}