// cmd/main.go
package main

import (
	"gin-template/controllers"
	"gin-template/infrastructure"
	"gin-template/usecases"

	"github.com/gin-gonic/gin"
)

func main() {
	// Initialize repository
	helloRepo := infrastructure.NewDbHelloRepository()

	// Initialize use case
	helloUseCase := usecases.NewHelloUseCase(helloRepo)

	// Initialize controller
	helloController := controllers.NewHelloController(helloUseCase)

	// Set up Gin router
	router := gin.Default()

	// Define routes
	router.GET("/hello", helloController.GetHello)

	// Start the server
	router.Run(":8083")
}
