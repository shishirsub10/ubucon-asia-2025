package main

import (
	"flag"
	"fmt"
	"os"
	"path/filepath"

	"github.com/gin-gonic/gin"
	"github.com/shishirsub10/ubucon-asia-2025/src/handlers"
)

// setupRouter sets up all routes
func setupRouter() *gin.Engine {
	router := gin.Default()

	// Serve static images from ~/animal_pictures at /images/
	homeDir, _ := os.UserHomeDir()
	imageDir := filepath.Join(homeDir, "animal_pictures")
	router.Static("/images", imageDir)

	router.GET("/", handlers.RootHandler)                // Gallery or list
	router.GET("/read", handlers.ReadFile)               // Read a file from ~/animal_pictures
	router.GET("/hello", handlers.HelloWorld)            // Simple hello endpoint
	router.POST("/ping", handlers.Ping)                  // Simple ping endpoint
	router.GET("/ping", handlers.PingForm)               // Serve ping form
	router.GET("/fetch", handlers.FetchForm)             // Fetch a URL
	router.POST("/fetch", handlers.FetchURL)             // Fetch a URL
	router.GET("/upload", handlers.UploadForm)           // Serve upload form
	router.POST("/upload", handlers.UploadAnimalPicture) // Upload to ~/animal_pictures

	return router
}

func main() {
	host := flag.String("host", "localhost", "Host to run the server on")
	port := flag.String("port", "8080", "Port to run the server on")
	flag.Parse()

	address := fmt.Sprintf("%s:%s", *host, *port)

	router := setupRouter()
	router.Run(address)
}
