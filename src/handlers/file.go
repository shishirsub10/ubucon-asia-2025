package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	utils "github.com/shishirsub10/ubucon-asia-2025/src/utils/file"
)

// ReadFile handles GET /read?file=<path>
func ReadFile(c *gin.Context) {
	filePath := c.Query("file")
	if filePath == "" {
		c.String(http.StatusBadRequest, "file query parameter is required")
		return
	}

	content, err := utils.ReadFileContent(filePath)
	if err != nil {
		c.String(http.StatusInternalServerError, "error reading file: %v", err)
		return
	}

	c.String(http.StatusOK, content)
}