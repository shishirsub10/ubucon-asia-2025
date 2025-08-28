package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	utils "github.com/shishirsub10/ubucon-asia-2025/src/utils/ping"
)

// Ping is the HTTP handler that calls the ping function
func Ping(c *gin.Context) {
	ip := c.PostForm("ip")
	if ip == "" {
		c.String(http.StatusBadRequest, "ip form parameter is required")
		return
	}

	output, err := utils.Ping(ip)
	if err != nil {
		c.String(http.StatusInternalServerError, "error executing ping: %v\nOutput: %s", err, output)
		return
	}

	c.String(http.StatusOK, output)
}

func PingForm(c *gin.Context) {
	c.Header("Content-Type", "text/html")
	c.String(http.StatusOK, `
		<!DOCTYPE html>
		<html>
		<head><title>Ping</title></head>
		<body>
			<h1>Ping Tool</h1>
			<form action="/ping" method="post">
				<label for="ip">IP address:</label>
				<input type="text" id="ip" name="ip" required>
				<button type="submit">Ping</button>
			</form>
		</body>
		</html>
	`)
}
