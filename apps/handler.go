package apps

import (
	"anaconda/utils"
	"net/http"

	"github.com/gin-gonic/gin"
)

type Handler struct {
	service *Service
}

func (h Handler) handleTest(c *gin.Context) {
	c.JSON(http.StatusOK, utils.Response("success"))
}
