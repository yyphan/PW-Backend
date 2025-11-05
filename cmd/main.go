package main

import (
	"log"

	"yyphan-pw/backend/internal/database"

	"github.com/gin-gonic/gin"
)

func main() {
	database.ConnectDatabase()

	router := gin.Default()
	router.SetTrustedProxies(nil)

	port := "8080"

	log.Printf("Server starting on port %s", port)
	router.Run("localhost:" + port)
}
