package main

import (
	"day9/songs"
	"github.com/gin-gonic/gin"
	"net/http"
)

func main() {
	r := gin.New()

	r.Use(gin.Recovery())
	r.GET("health", healthCheck)
	r.POST("/song", songs.CreateHandler)
	r.GET("/song/:id", songs.GetHandler)
	r.DELETE("/song/:id", songs.DeleteHandler)
	r.PUT("/song/:id", songs.UpdateHandler)
	r.Run("localhost:8080")
}

func healthCheck(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{"status": "Healthy"})
}
