package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// HelloWorld handles GET /hello
func HelloWorld(c *gin.Context) {
	c.String(http.StatusOK, "hello world")
}
