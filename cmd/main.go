package main

import (
	"log"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	router.SetTrustedProxies(nil)

	port := "8080"

	log.Printf("Server starting on port %s", port)
	router.Run("localhost:" + port)
}
