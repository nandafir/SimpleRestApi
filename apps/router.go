package apps

import (
	"github.com/gin-gonic/gin"
)

// Register register routes
func Register(router *gin.Engine, service *Service) {

	handler := &Handler{
		service: service,
	}
	v1 := router.Group("/v1/local/")
	{
		v1.GET("", handler.handleTest)
	}
}
