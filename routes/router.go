package routes

import (
	"github.com/Rodrigosiq03/serverless-go-api/controllers"
	"github.com/gin-gonic/gin"
)

func InitRoutes(r *gin.Engine) {
	r.GET("/get-all-products", controllers.GetAllProducts)
}
