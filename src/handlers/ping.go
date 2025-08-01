package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	utils "github.com/shishirsub10/ubucon-asia-2025/src/utils/ping"
)

// Ping is the HTTP handler that calls the ping function
func Ping(c *gin.Context) {
	ip := c.Query("ip")
	if ip == "" {
		c.String(http.StatusBadRequest, "ip query parameter is required")
		return
	}

	output, err := utils.Ping(ip)
	if err != nil {
		c.String(http.StatusInternalServerError, "error executing ping: %v\nOutput: %s", err, output)
		return
	}

	c.String(http.StatusOK, output)
}
