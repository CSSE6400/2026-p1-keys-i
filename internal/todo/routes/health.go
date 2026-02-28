package routes

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func RegisterHealth(rtr *gin.RouterGroup) {
	rtr.GET("/health", func(ctx *gin.Context) {
		ctx.IndentedJSON(http.StatusOK, gin.H{"status": "ok"})
	})
}
