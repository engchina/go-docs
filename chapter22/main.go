package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func main() {
	r := gin.Default()
	r.GET("/ping", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, gin.H{"msg": "Pong"})
	})
	err := r.Run("127.0.0.1:8181")
	if err != nil {
		return
	}
}
