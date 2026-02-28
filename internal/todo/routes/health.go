package routes

import "github.com/gin-gonic/gin"

func RegisterHealth(rtr *gin.RouterGroup) {
	rtr.GET("/health", func(ctx *gin.Context) {
		ctx.JSON(200, gin.H{"status": "ok"})
	})
}
