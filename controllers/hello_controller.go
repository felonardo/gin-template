// controllers/hello_controller.go
package controllers

import (
	"gin-template/usecases"
	"net/http"

	"github.com/gin-gonic/gin"
)

type HelloController struct {
	helloUseCase usecases.HelloUseCase
}

func NewHelloController(u usecases.HelloUseCase) *HelloController {
	return &HelloController{
		helloUseCase: u,
	}
}

func (ctrl *HelloController) GetHello(c *gin.Context) {
	helloMessage := ctrl.helloUseCase.GetHelloMessage()
	c.JSON(http.StatusOK, gin.H{"message": helloMessage.Message})
}
