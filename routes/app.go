package routes

import "github.com/gin-gonic/gin"

func App() {
	router := gin.Default()

	InitRoutes(router)

	router.Run()
}
