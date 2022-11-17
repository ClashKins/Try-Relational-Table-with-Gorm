package routers

import (
	"LATIHAN1/controllers"

	"github.com/gin-gonic/gin"
)

func StartServer() *gin.Engine{
	router := gin.Default()

	router.GET("/marts/:id", controllers.GetOneProducts)
	router.GET("/marts", controllers.GetAllProducts)
	router.POST("/marts", controllers.CreateMarts)
	router.PATCH("/marts/:id", controllers.UpdateMarts)
	router.DELETE("/marts/:id", controllers.DeleteMart)

	return router
}