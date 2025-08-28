package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	utils "github.com/shishirsub10/ubucon-asia-2025/src/utils/http_request"
)

// FetchURL takes a user-supplied URL and returns its content.
// Example: POST /fetch with form field "url=http://example.com"
//
// ⚠️ Deliberately Insecure (for demo purposes):
//   - No validation or sanitization of the input URL.
//   - Allows SSRF (Server-Side Request Forgery) attacks, e.g. fetching
//     http://localhost:8080/secret or internal metadata services.
//
// This insecure handler is useful for demonstrating how AppArmor
// can be used to restrict network access even when the application code is vulnerable.
func FetchURL(c *gin.Context) {
	url := c.PostForm("url")
	if url == "" {
		c.String(http.StatusBadRequest, "url form parameter is required")
		return
	}

	content, err := utils.FetchURLContent(url)
	if err != nil {
		c.String(http.StatusInternalServerError, "error fetching url: %v", err)
		return
	}

	c.String(http.StatusOK, content)
}

// FetchForm serves a simple HTML form for testing the FetchURL POST endpoint.
func FetchForm(c *gin.Context) {
	const page = `<!DOCTYPE html>
<html lang="en">
<head>
  <meta charset="utf-8">
  <title>Fetch URL (Insecure Demo)</title>
</head>
<body style="font-family: sans-serif; max-width: 640px; margin: 40px auto;">
  <h1>🌐 Fetch URL (Deliberately Insecure)</h1>
  <form method="POST" action="/fetch">
    <p>
      <label for="url">URL to fetch:</label><br>
      <input type="text" name="url" id="url" style="width:100%;" placeholder="http://example.com" required>
    </p>
    <p>
      <button type="submit">Fetch</button>
      <a href="/" style="margin-left: 8px;">Back to Gallery</a>
    </p>
  </form>
</body>
</html>`
	c.Data(http.StatusOK, "text/html; charset=utf-8", []byte(page))
}
