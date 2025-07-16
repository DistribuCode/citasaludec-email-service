package routes

import (
    "github.com/gin-gonic/gin"
    "email-service/internal/controllers"
)

func RegisterRoutes(r *gin.Engine) {
    r.POST("/email", controllers.SendEmail)
}
