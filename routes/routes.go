package routes

import (
	"github.com/gin-gonic/gin"
	"messenger-backend/controllers"
)

func RegisterRoutes(router *gin.Engine) {
	userGroup := router.Group("/user")
	{
		userGroup.POST("/validate", controllers.Validate)
		userGroup.POST("/register", controllers.Register)
		userGroup.POST("/login", controllers.Login)
		userGroup.GET("/profile", controllers.Profile)
		userGroup.POST("/send-verification", controllers.SendVerificationCode)
		userGroup.POST("/verify-code", controllers.VerifyCode)
	}
}