package routes

import (
	"github.com/Maliud/GO_JWT_Authentication_with_Gin_Gonic/controllers"
	"github.com/gin-gonic/gin"
)

func AuthRoutes(incomingRoutes *gin.Engine){
	incomingRoutes.POST("users/signup",controllers.Signup())
	incomingRoutes.POST("users/login", controllers.Login())
}