package todo

import (
	"github.com/CSSE6400/2026-p1-keys-i/internal/todo/routes"
	"github.com/gin-gonic/gin"
)

func NewRouter() *gin.Engine {
	rtr := gin.New()
	rtr.Use(gin.Recovery())

	// route rego with version group
	v1 := rtr.Group("/api/v1")
	{
		routes.RegisterHealth(v1)
		routes.RegisterTODO(v1)
	}

	return rtr
}
