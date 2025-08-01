package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	utils "github.com/shishirsub10/ubucon-asia-2025/src/utils/http_request"
)

func FetchURL(c *gin.Context) {
	url := c.Query("url")
	if url == "" {
		c.String(http.StatusBadRequest, "url query parameter is required")
		return
	}

	content, err := utils.FetchURLContent(url)
	if err != nil {
		c.String(http.StatusInternalServerError, "error fetching url: %v", err)
		return
	}

	c.String(http.StatusOK, content)
}