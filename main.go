package main

import (
	"flag"
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/shishirsub10/ubucon-asia-2025/src/handlers"
)

// setupRouter sets up all routes
func setupRouter() *gin.Engine {
	router := gin.Default()
	router.GET("/read", handlers.ReadFile)
	router.GET("/hello", handlers.HelloWorld)
	router.GET("/ping", handlers.Ping)
	router.GET("/fetch", handlers.FetchURL)
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

