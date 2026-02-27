package main

import (
	"log"
	"os"

	"yyphan-pw/backend/internal/database"
	"yyphan-pw/backend/internal/routers"

	"github.com/gin-gonic/gin"
)

func main() {
	database.ConnectDatabase()

	router := gin.Default()
	router.SetTrustedProxies(nil)

	routers.InitRouter(router)

	port := os.Getenv("APP_PORT")

	if port == "" {
		port = "8080"
	}

	log.Printf("Server starting on port %s", port)
	router.Run(":" + port)
}
