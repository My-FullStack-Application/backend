package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func heartbeat(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{
		"message": "OK",
	})
}

func main() {
	server := gin.Default()
	server.GET("/heartbeat", heartbeat)
	server.Run("localhost:8080")
}
