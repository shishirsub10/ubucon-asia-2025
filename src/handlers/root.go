package handlers

import (
	"html/template"
	"net/http"
	"os"
	"path/filepath"
	"strings"

	"github.com/gin-gonic/gin"
)

// RootHandler ensures ~/animal_pictures exists and shows an HTML gallery of images.
func RootHandler(c *gin.Context) {
	homeDir, err := os.UserHomeDir()
	if err != nil {
		c.String(http.StatusInternalServerError, "cannot determine home directory")
		return
	}
	baseDir := filepath.Join(homeDir, "animal_pictures")

	// Create directory if it doesn't exist
	if _, err := os.Stat(baseDir); os.IsNotExist(err) {
		if err := os.MkdirAll(baseDir, 0o755); err != nil {
			c.String(http.StatusInternalServerError, "failed to create animal_pictures directory")
			return
		}
	}

	entries, err := os.ReadDir(baseDir)
	if err != nil {
		c.String(http.StatusInternalServerError, "failed to read animal_pictures directory")
		return
	}

	// Collect only picture files (jpg, jpeg, png, gif)
	var pictures []string
	for _, e := range entries {
		if e.IsDir() {
			continue
		}
		ext := strings.ToLower(filepath.Ext(e.Name()))
		if ext == ".jpg" || ext == ".jpeg" || ext == ".png" || ext == ".gif" {
			pictures = append(pictures, "/images/"+e.Name())
		}
	}

	// Serve HTML gallery, even if there are no images
	tpl := `
	<!DOCTYPE html>
	<html>
	<head><title>Animal Gallery</title></head>
	<body>
		<h1>🐾 Animal Pictures</h1>
		<nav>
			<ul>
				<li><a href="/">Gallery (Home)</a></li>
				<li><a href="/upload">Upload Picture</a></li>
				<li><a href="/hello">Hello</a></li>
				<li><a href="/ping">Ping</a></li>
				<li><a href="/fetch">Fetch URL</a></li>
			</ul>
		</nav>
		{{if .}}
			{{range .}}
				<div style="display:inline-block; margin:10px;">
					<img src="{{.}}" style="max-width:200px; max-height:200px;" />
				</div>
			{{end}}
		{{else}}
			<p>No pictures found. <a href="/upload">Upload one?</a></p>
		{{end}}
	</body>
	</html>`
	t, _ := template.New("gallery").Parse(tpl)
	c.Header("Content-Type", "text/html; charset=utf-8")
	c.Status(http.StatusOK)
	t.Execute(c.Writer, pictures)
}
