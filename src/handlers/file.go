package handlers

import (
	"net/http"
	"os"
	"path/filepath"

	"github.com/gin-gonic/gin"
)

// ReadFile reads a file from ~/animal_pictures based on the `filename` query parameter.
// Example: GET /read?filename=dog.png
//
// ⚠️ Deliberately Insecure (for demo purposes):
// - No validation or sanitization is done on `filename`.
// - An attacker can supply values like `../secret.txt` to attempt directory traversal.
// - This allows reading files outside ~/animal_pictures if not confined by AppArmor.
//
// In a real-world application you should validate the input path,
// ensure it stays inside the intended directory, and reject traversal attempts.
func ReadFile(c *gin.Context) {
	filename := c.Query("filename")
	if filename == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Filename is required"})
		return
	}

	homeDir, _ := os.UserHomeDir()
	filePath := filepath.Join(homeDir, "animal_pictures", filename)

	// Directly serve the requested file to the client.
	// No error handling here for missing/invalid files (also insecure).
	c.File(filePath)
}

// UploadAnimalPicture lets the client control the destination path via "dest".
// This is intentionally insecure for demo purposes.
func UploadAnimalPicture(c *gin.Context) {
	file, err := c.FormFile("animal_picture")
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "No file received"})
		return
	}

	// Let the attacker (you, for the demo) specify where to save:
	// e.g. dest=../../../../tmp/cat4.jpeg
	dest := c.PostForm("dest")
	if dest == "" {
		// fallback to original filename if no dest provided
		dest = file.Filename
	}

	homeDir, _ := os.UserHomeDir()

	// ❌ Intentionally vulnerable: user-controlled "dest" is joined with base,
	// enabling traversal outside ~/animal_pictures via ../ segments.
	base := filepath.Join(homeDir, "animal_pictures")
	savePath := filepath.Join(base, dest) // filepath.Join will normalize .. and escape base

	// Create parent dirs and save
	if err := os.MkdirAll(filepath.Dir(savePath), 0o755); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create directory"})
		return
	}
	if err := c.SaveUploadedFile(file, savePath); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to save file"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "File uploaded", "path": savePath})
}

// UploadForm serves an HTML form that allows picking a file and specifying "dest".
func UploadForm(c *gin.Context) {
	const page = `<!DOCTYPE html>
<html lang="en">
<head><meta charset="utf-8"><title>Upload (Insecure Demo)</title></head>
<body style="font-family: sans-serif; max-width: 640px; margin: 40px auto;">
  <h1>📤 Upload (Deliberately Insecure)</h1>
  <form method="POST" action="/upload" enctype="multipart/form-data">
    <p><input type="file" name="animal_picture" required></p>
    <p>
      <label>dest (e.g. cat4.jpeg):</label><br>
      <input type="text" name="dest" style="width: 100%;" placeholder="my-cat.jpeg">
    </p>
    <p>
      <button type="submit">Upload</button>
      <a href="/" style="margin-left: 8px;">Back to Gallery</a>
    </p>
  </form>
</body>
</html>`
	c.Data(http.StatusOK, "text/html; charset=utf-8", []byte(page))
}
