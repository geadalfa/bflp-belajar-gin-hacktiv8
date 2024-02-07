package routers

import (
	"belajar-gin/controllers"

	"github.com/gin-gonic/gin"

	_ "belajar-gin/docs"

	ginSwagger "github.com/swaggo/gin-swagger"

	swaggerfiles "github.com/swaggo/files"
)

// @title Car API
// @version 1.0
// @description This is a sample service for managing cars
// @termsOfService: https://swagger.io/terms/
// @contact.name API Support
// @contact.email soberkoder@swagger.io
// @license.name Apache 2.0
// @license.url: http://www.apache.org/licenses/LICENSE-2.0.html
// @host localhost:8080
// @BasePath /
func StartServer() *gin.Engine {
	router := gin.Default()
	// Create
	router.POST("/cars", controllers.CreateCar)
	// Update
	router.PUT("/cars/:id", controllers.UpdateCar)
	// Read
	router.GET("/cars/:id", controllers.GetCar)
	// Delete
	router.DELETE("/cars/:id", controllers.DeleteCar)
	// Read All
	router.GET("/cars", controllers.GetAllCar)

	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))

	return router
}
