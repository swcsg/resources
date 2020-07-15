package healthz

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// Healthz 健康检查
func Healthz(c *gin.Context) {
	c.String(http.StatusOK, "")
	return
}